package main

import (
	"context"
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego-component/eetcd"
	"github.com/gotomicro/ego-component/eetcd/examples/helloworld"
	"github.com/gotomicro/ego-component/eetcd/registry"
	"github.com/gotomicro/ego/client/egrpc"
	"github.com/gotomicro/ego/core/elog"
	"log"
)

// export EGO_DEBUG=true
func main() {
	if err := ego.New().Invoker(
		invokerGrpc,
		callGrpc,
	).Run(); err != nil {
		elog.Error("startup", elog.FieldErr(err))
	}
}

var grpcComp helloworld.GreeterClient

func invokerGrpc() error {
	registry.Load("registry").Build(registry.WithClientEtcd(eetcd.Load("etcd").Build()))
	grpcConn := egrpc.Load("grpc.test").Build()
	grpcComp = helloworld.NewGreeterClient(grpcConn.ClientConn)
	return nil
}

func callGrpc() error {
	log.Println("-------1")
	_, err := grpcComp.SayHello(context.Background(), &helloworld.HelloRequest{
		Name: "i am client",
	})
	log.Println("-------2")
	if err != nil {
		log.Println("-------4")
		return err
	}
	log.Println("-------3")

	_, err = grpcComp.SayHello(context.Background(), &helloworld.HelloRequest{
		Name: "error2",
	})
	if err != nil {
		return err
	}
	return nil
}