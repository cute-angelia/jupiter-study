package helloworld

import (
	"context"
	"github.com/douyu/jupiter/pkg/server/xgrpc"
	"jupiter-study/proto/helloworld"
)

type GreeterServer struct {
	Server *xgrpc.Server
}

func (server GreeterServer) SayHello(context context.Context, request *helloworld.HelloRequest) (response *helloworld.HelloReply, err error) {
	rsp := helloworld.HelloReply{}
	rsp.Message = request.Name + " good"
	response = &rsp
	return
}
