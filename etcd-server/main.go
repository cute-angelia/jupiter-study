package main

import (
	"github.com/douyu/jupiter"
	"github.com/douyu/jupiter/pkg"
	compound_registry "github.com/douyu/jupiter/pkg/registry/compound"
	etcdv3_registry "github.com/douyu/jupiter/pkg/registry/etcdv3"
	"github.com/douyu/jupiter/pkg/server/xgrpc"
	"github.com/douyu/jupiter/pkg/xlog"
	hellword2 "jupiter-study/etcd-server/internel/grpc/helloword"
	"jupiter-study/library/config"
	"jupiter-study/proto/helloworld"
)

func main() {
	// init config
	config.NewConfig().InitConfig("study")

	eng := NewEngine()
	eng.SetRegistry(
		compound_registry.New(
			etcdv3_registry.StdConfig("hz").Build(),
		),
	)
	//eng.SetGovernor("0.0.0.0:0")
	if err := eng.Run(); err != nil {
		xlog.Error(err.Error())
	}
}

type Engine struct {
	jupiter.Application
}

func NewEngine() *Engine {
	eng := &Engine{}
	if err := eng.Startup(
		eng.serveGRPC,
	); err != nil {
		xlog.Panic("startup", xlog.Any("err", err))
	}
	return eng
}


func (eng *Engine) serveGRPC() error {
	pkg.SetName("etcd-server")
	pkg.SetAppID("appid 2323")

	grpc := xgrpc.StdConfig("grpc").Build()

	// register
	helloworld.RegisterGreeterServer(grpc.Server, &hellword2.GreeterServer{Server: grpc})

	return eng.Serve(grpc)
}
