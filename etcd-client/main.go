package main

import (
	"context"
	"fmt"
	"jupiter-study/library/config"
	"log"
	"time"

	"github.com/douyu/jupiter"
	"github.com/douyu/jupiter/pkg/client/grpc"
	"github.com/douyu/jupiter/pkg/client/grpc/balancer"
	"github.com/douyu/jupiter/pkg/client/grpc/resolver"
	"github.com/douyu/jupiter/pkg/registry/etcdv3"
	"github.com/douyu/jupiter/pkg/xlog"

	"jupiter-study/proto/goodbye"
	"jupiter-study/proto/helloworld"
)

func main() {
	// init config
	config.NewConfig().InitConfig("study")

	eng := NewEngine()
	if err := eng.Run(); err != nil {
		xlog.Error(err.Error())
	}
	fmt.Printf("111 = %+v\n", 111)
}

type Engine struct {
	jupiter.Application
}

func NewEngine() *Engine {
	eng := &Engine{}
	if err := eng.Startup(
		eng.initResolver,
		eng.consumer,
	); err != nil {
		xlog.Panic("startup", xlog.Any("err", err))
	}
	return eng
}

func (eng *Engine) initResolver() error {
	resolver.Register("etcd", etcdv3.StdConfig("hz").Build())
	return nil
}

func (eng *Engine) consumer() error {
	config := grpc.StdConfig("etcd-server")
	config.Debug = true
	config.BalancerName = balancer.NameSmoothWeightRoundRobin
	client := helloworld.NewGreeterClient(config.Build())

	config2 := grpc.StdConfig("etcd-server-2")
	client2 := goodbye.NewGoodByeClient(config2.Build())

	go func() {
		for {
			if resp2, err := client2.SayGoodBye(context.Background(), &goodbye.SayGoodByeReq{
				Name: "jupiter",
			}); err != nil {
				log.Println(err)
			} else {
				log.Println(resp2.Message)
			}

			resp, err := client.SayHello(context.Background(), &helloworld.HelloRequest{
				Name: "jupiter",
			})
			if err != nil {
				fmt.Printf("err = %+v\n", err)
				xlog.Error(err.Error())
			} else {
				fmt.Printf("resp.Message = %+v\n", resp.Message)
				xlog.Info("receive response", xlog.String("resp", resp.Message))
			}
			time.Sleep(3 * time.Second)
		}
	}()
	return nil
}
