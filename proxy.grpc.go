package main

import (
	"context"
	"errors"
	"sync/atomic"

	// "fmt"
	"io"
	"strconv"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

func redirectGRPC(ss grpc.ServerStream, conn *grpc.ClientConn, custom map[string][]string) error {
	fullMethodName, ok := grpc.MethodFromServerStream(ss)
	if !ok {
		return status.Errorf(codes.Internal, "failed to get method from server stream")
	}

	ssctx := ss.Context()
	md, _ := metadata.FromIncomingContext(ssctx)
	for name, h := range custom { // 注入自定义头部
		md.Set(name, h...)
	}
	outCtx := metadata.NewOutgoingContext(ssctx, md.Copy())
	redirectTo, err := grpc.NewClientStream(outCtx, &grpc.StreamDesc{
		ServerStreams: true,
		ClientStreams: true,
	}, conn, fullMethodName)
	if err != nil {
		return status.Errorf(codes.Canceled, "the ClientConn may has been closed")
	}
	ErrChanA2B := redirectGRPC2server(ss, redirectTo) // req
	ErrChanB2A := redirectGRPC2client(redirectTo, ss) // resp
	for i := 0; i < 2; i++ {
	L:
		select {
		case err = <-ErrChanA2B:
			if err == io.EOF {
				// success
				_ = redirectTo.CloseSend()
				break L
			} else {
				return status.Errorf(codes.Internal, "failed proxying s2c: %v", err)
			}
		case err = <-ErrChanB2A:
			ss.SetTrailer(redirectTo.Trailer())
			if err != io.EOF {
				return err
			}
			return nil
		}
	}
	return status.Errorf(codes.Internal, "should never reach this stage.")
}

func redirectGRPC2server(src grpc.ServerStream, dst grpc.ClientStream) chan error {
	ret := make(chan error, 1)
	go func() {
		f := &anypb.Any{}
		for i := 0; ; i++ {
			if err := src.RecvMsg(f); err != nil {
				ret <- err // this can be io.EOF which is happy case
				break
			}
			if err := dst.SendMsg(f); err != nil {
				ret <- err
				break
			}
		}
	}()
	return ret
}

func redirectGRPC2client(src grpc.ClientStream, dst grpc.ServerStream) chan error {
	ret := make(chan error, 1)
	go func() {
		f := &anypb.Any{}
		if err := src.RecvMsg(f); err != nil {
			ret <- err // this can be io.EOF which is happy case
			return
		}
		md, err := src.Header()
		if err != nil {
			ret <- err
			return
		}
		if err := dst.SendHeader(md); err != nil {
			ret <- err
			return
		}

		if err := dst.SendMsg(f); err != nil {
			ret <- err
			return
		}
		for {
			if err := src.RecvMsg(f); err != nil {
				ret <- err
				break
			}
			if err := dst.SendMsg(f); err != nil {
				ret <- err
				break
			}
		}
	}()
	return ret
}

func handlerInGRPC(nodeActiveContext context.Context) func(_ interface{}, s grpc.ServerStream) (err error) {
	return func(_ interface{}, s grpc.ServerStream) (err error) {
		md, ok := metadata.FromIncomingContext(s.Context())
		trace_id := "UNKNOWN"
		from := "UNKNOWN"
		if ok {
			trace_ids := md.Get("trace_id")
			if len(trace_ids) > 0 {
				trace_id = trace_ids[0]
			}
			froms := md.Get("from")
			if len(froms) > 0 {
				from = froms[0]
			}
		}
		srv_name, mth := fullServiceNameFrom(s)
		log_f := logrus.Fields{"direction": "IN", "want_srv": srv_name, "want_mth": mth, "trace_id": trace_id, "from": from}
		logger.WithFields(log_f).Info()

		if _, ok := CONFIG.AllowService["*"]; !ok {
			if _, ok := CONFIG.AllowService[from]; !ok {
				logger.WithFields(log_f).Warn("not allowed service")
				return errors.New("not allowed service")
			}
		}

		conn, _, err := P.local.Conn(nodeActiveContext)
		if err != nil {
			logger.WithFields(log_f).Error(err)
			return err
		}
		err = redirectGRPC(s, conn, nil)
		if err != nil {
			logger.WithFields(log_f).Error(err)
			return err
		}
		return nil
	}
}

var reqid int64 = 1

func makereqid() int {
	if reqid > 1000 {
		atomic.StoreInt64(&reqid, 1)
	} else {
		atomic.AddInt64(&reqid, 1)
	}
	return int(reqid)
}

func handlerOutGRPC(nodeActiveContext context.Context) func(_ interface{}, s grpc.ServerStream) (err error) {
	return func(_ interface{}, s grpc.ServerStream) (err error) {
		md, ok := metadata.FromIncomingContext(s.Context())
		var trace string
		reqid := int(makereqid())
		if ok {
			trace_ids := md.Get("trace_id")
			if len(trace_ids) > 0 {
				trace = trace_ids[0]
				goto HaveTrace
			}
		}
		trace = strconv.Itoa(reqid + int(node_id))
	HaveTrace:
		srv_name, mth := fullServiceNameFrom(s)
		log_f := logrus.Fields{"direction": "OUT", "want_srv": srv_name, "want_mth": mth, "trace_id": trace}
		logger.WithFields(log_f).Info()

		endpoint := P.discoverer.direct(srv_name)
		if endpoint == nil {
			logger.WithFields(log_f).Error("couldnt find lived node")
			return errors.New("couldnt find " + srv_name)
		}

		conn, delay, err := (*nodeGRPC)(endpoint).Conn(nodeActiveContext)
		if err != nil {
			logger.WithFields(log_f).Error("get redirect_to conn failed", err)
			endpoint.rm()
			return errors.New("get redirect_to conn failed")
		}

		err = redirectGRPC(s, conn, map[string][]string{
			"trace_id": {trace},
			"from":     {string(name())},
		})
		if err != nil {
			logger.WithFields(log_f).Error("redirect failed", err)
			endpoint.rm()
			return err
		}
		endpoint.SetDelay(delay)
		return nil
	}
}
