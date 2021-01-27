package main

import (
	"github.com/douyu/jupiter"
	"github.com/douyu/jupiter/pkg"
	compound_registry "github.com/douyu/jupiter/pkg/registry/compound"
	etcdv3_registry "github.com/douyu/jupiter/pkg/registry/etcdv3"
	"github.com/douyu/jupiter/pkg/server"
	"github.com/douyu/jupiter/pkg/server/xgrpc"
	"github.com/douyu/jupiter/pkg/xlog"
	gGoodBye "jupiter-study/etcd-server-2/internel/grpc/goodbye"
	"jupiter-study/library/config"
	"jupiter-study/proto/goodbye"
	"log"
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
	server := xgrpc.StdConfig("grpc").Build()

	z := server.Info()
	z.Name = "srv-demo-server"

	log.Println("xcdadfs", server.Info(), server.Info().Name, server.Info().Label(), z.Name)

	// register
	goodbye.RegisterGoodByeServer(server.Server, &gGoodBye.GoodbyeServer{Server: server})

	return eng.Serve(server)
}

func defaultServiceInfo() *server.ServiceInfo {
	si := server.ServiceInfo{
		Name:       pkg.Name(),
		AppID:      pkg.AppID(),
		Weight:     100,
		Enable:     true,
		Healthy:    true,
		Metadata:   make(map[string]string),
		Region:     pkg.AppRegion(),
		Zone:       pkg.AppZone(),
		Kind:       0,
		Deployment: "",
		Group:      "",
	}
	si.Metadata["appMode"] = pkg.AppMode()
	si.Metadata["appHost"] = pkg.AppHost()
	si.Metadata["startTime"] = pkg.StartTime()
	si.Metadata["buildTime"] = pkg.BuildTime()
	si.Metadata["appVersion"] = pkg.AppVersion()
	si.Metadata["jupiterVersion"] = pkg.JupiterVersion()
	return &si
}
