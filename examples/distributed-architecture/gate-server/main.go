package main

import (
	"github.com/dobyte/due"
	cluster "github.com/dobyte/due/cluster/gate"
	"github.com/dobyte/due/locate/redis"
	"github.com/dobyte/due/network/tcp"
	"github.com/dobyte/due/registry/etcd"
	"github.com/dobyte/due/transport/grpc"
)

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建网关组件
	gate := cluster.NewGate(
		cluster.WithServer(tcp.NewServer()),
		cluster.WithLocator(redis.NewLocator()),
		cluster.WithRegistry(etcd.NewRegistry()),
		cluster.WithTransporter(grpc.NewTransporter()),
	)
	// 添加网关组件
	container.Add(gate)
	// 启动容器
	container.Serve()
}
