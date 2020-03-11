package proxy

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"strings"
)

type StreamDirector func(ctx context.Context, serviceAndMethodName string) (context.Context, *grpc.ClientConn, error)

func ReadNames(serviceAndMethodName string)(string,string,bool){
	t := strings.Split(serviceAndMethodName,"/")
	t2 := strings.Split(t[1], ".")[0]
	if len(t) != 3 {
		return "","",false
	}
	return t2, t[2], true
}
