package main

import (
	"context"
	"github.com/dobyte/due"
	cluster "github.com/dobyte/due/cluster/mesh"
	"github.com/dobyte/due/locate/redis"
	"github.com/dobyte/due/log"
	"github.com/dobyte/due/mode"
	"github.com/dobyte/due/registry/consul"
	"github.com/dobyte/due/transport/rpcx"
)

func main() {
	// 开启调试模式
	mode.SetMode(mode.DebugMode)
	// 创建容器
	container := due.NewContainer()
	// 创建网格组件
	mesh := cluster.NewMesh(
		cluster.WithLocator(redis.NewLocator()),
		cluster.WithRegistry(consul.NewRegistry()),
		cluster.WithTransporter(rpcx.NewTransporter()),
	)
	// 初始化业务
	NewWalletService(mesh.Proxy()).Init()
	// 添加网格组件
	container.Add(mesh)
	// 启动容器
	container.Serve()
}

// WalletService 钱包服务
type WalletService struct {
	proxy *cluster.Proxy
}

type IncrGoldRequest struct {
	UID  int64
	Gold int64
}

type IncrGoldReply struct {
}

func NewWalletService(proxy *cluster.Proxy) *WalletService {
	return &WalletService{proxy: proxy}
}

func (w *WalletService) Init() {
	w.proxy.AddServiceProvider("wallet", "Wallet", w)
}

func (w *WalletService) IncrGold(ctx context.Context, req *IncrGoldRequest, reply *IncrGoldReply) error {
	log.Infof("incr %d gold success for uid: %d", req.Gold, req.UID)

	return nil
}
