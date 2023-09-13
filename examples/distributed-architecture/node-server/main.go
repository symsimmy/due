package main

import (
	"github.com/dobyte/due"
	cluster "github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/locate/redis"
	"github.com/dobyte/due/registry/consul"
	"github.com/dobyte/due/transport/grpc"
)

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建节点组件
	node := cluster.NewNode(
		cluster.WithLocator(redis.NewLocator()),
		cluster.WithRegistry(consul.NewRegistry()),
		cluster.WithTransporter(grpc.NewTransporter()),
	)
	// 注册路由
	node.Proxy().Router().AddRouteHandler(1, false, greetHandler)
	// 添加组件
	container.Add(node)
	// 启动服务器
	container.Serve()
}

func greetHandler(ctx *cluster.Context) {
	_ = ctx.Response([]byte("hello world~~"))
}
