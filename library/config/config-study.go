package config

var configStudy = `
[jupiter.server.grpc]
	port = 0

[jupiter.registry.hz]
    configKey = "jupiter.etcdv3.default"
    connectTimeout = "1s"
    endpoints=["127.0.0.1:2379"]
    secure = false

[jupiter.etcdv3.default]
    endpoints = ["127.0.0.1:2379"]
    secure = false

[jupiter.client.etcd-server]
    address = "etcd:///etcd-server"
    balancerName = "round_robin"
    block =  false
    dialTimeout = "0s"

[jupiter.client.etcd-server-2]
    address = "etcd:///etcd-server-2"
    balancerName = "round_robin"
    block =  false
    dialTimeout = "0s"

`
