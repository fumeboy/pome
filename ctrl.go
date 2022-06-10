package main

import (
	"bufio"
	context "context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"pome/ctrl"
	"strconv"
)

type Server struct{}

var SidecarSrv ctrl.SidecarServer = &Server{}

func (s *Server) SearchLogByTraceID(ctx context.Context, request *ctrl.SearchLogByTraceIDRequest) (e *ctrl.SearchLogByTraceIDResponse, e2 error) {
	pwd, _ := os.Getwd()
	files, _ := ioutil.ReadDir(pwd)
	file_time_list := []string{}
	for _, file := range files {
		if file.Name()[:7] == "log.txt" && len(file.Name()) > 7 {
			time, _ := strconv.Atoi(file.Name()[8:])
			if request.TimeStart <= int64(time) && (request.TimeEnd == 0 || request.TimeEnd >= int64(time)) {
				file_time_list = append(file_time_list, file.Name()[8:])
			}
		}
	}
	logs := []string{}
	for _, filetime := range file_time_list {
		f, e := os.Open("log.txt." + filetime)
		if e != nil {
			continue
		}
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			line := buf.Text()
			fmt.Println(line)
			jmap := map[string]string{}
			json.Unmarshal([]byte(line), &jmap)
			fmt.Println(jmap)
			if trace_id, ok := jmap["trace_id"]; ok && trace_id == strconv.Itoa(int(request.TraceId)) {
				logs = append(logs, line)
			}
		}
	}
	return &ctrl.SearchLogByTraceIDResponse{
		LogRecords: logs,
	}, nil
}

func (s *Server) Stop(ctx context.Context, request *ctrl.StopReq) (resp *ctrl.StopResp, err error) {
	if request.NodeId == node_id {
		P.stop = true
		P.cancel()
		return &ctrl.StopResp{
			Status: 1,
			Msg:    "",
		}, nil
	}
	return &ctrl.StopResp{
		Status: 2,
		Msg:    "bad node_id",
	}, nil
}

func (s *Server) Start(ctx context.Context, request *ctrl.StartReq) (resp *ctrl.StartResp, err error) {
	if P.stop == false {
		return &ctrl.StartResp{
			Status: 2,
			Msg:    "have started",
		}, nil
	}
	P.startCh <- true
	return &ctrl.StartResp{
		Status: 1,
		Msg:    "",
	}, nil
}

func (s *Server) ReadConfig(ctx context.Context, request *ctrl.RCReq) (resp *ctrl.RCResp, err error) {
	c := ReadConfig()
	return &ctrl.RCResp{Context: c}, nil
}

func (s *Server) UpdateConfig(ctx context.Context, request *ctrl.UCReq) (resp *ctrl.UCResp, err error) {
	e := UpdateConfig([]byte(request.Context))
	if e == nil {
		return &ctrl.UCResp{Status: 1, Msg: ""}, nil
	}
	return &ctrl.UCResp{Status: 2, Msg: e.Error()}, nil
}
