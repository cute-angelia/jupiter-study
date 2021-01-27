package goodbye

import (
	"context"
	"github.com/douyu/jupiter/pkg/server/xgrpc"
	"jupiter-study/proto/goodbye"
)

type GoodbyeServer struct {
	Server *xgrpc.Server
}

func (server GoodbyeServer) SayGoodBye (context context.Context, request *goodbye.SayGoodByeReq) (response *goodbye.SayGoodByeResp, err error) {
	rsp := goodbye.SayGoodByeResp{}
	rsp.Message = request.Name + " goodbye"
	response = &rsp
	return
}
