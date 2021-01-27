package config

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	_ "github.com/douyu/jupiter/pkg/datasource/etcdv3"
)

type config struct {
}

func NewConfig() *config {
	return &config{
	}
}

func (e *config) InitConfig(project string) {
	etcdCfg := clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	}
	cli, _ := clientv3.New(etcdCfg)
	cli.Put(context.Background(), fmt.Sprintf("config-%s", project), configStudy)
}
