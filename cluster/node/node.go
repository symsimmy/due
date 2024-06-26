package node

import (
	"context"
	"fmt"
	"github.com/symsimmy/due/cluster"
	"github.com/symsimmy/due/component"
	"github.com/symsimmy/due/log"
	"github.com/symsimmy/due/registry"
	"github.com/symsimmy/due/transport"
	"github.com/symsimmy/due/utils/xcall"
	"strconv"
	"sync/atomic"
	"time"
	"unsafe"
)

type Node struct {
	component.Base
	opts     *options
	ctx      context.Context
	cancel   context.CancelFunc
	state    cluster.State
	events   *Events
	router   *Router
	proxy    *Proxy
	instance *registry.ServiceInstance
	rpc      transport.Server
	fnChan   chan func()
}

func NewNode(opts ...Option) *Node {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	n := &Node{}
	n.opts = o
	n.events = newEvents(n)
	n.router = newRouter(n)
	n.proxy = newProxy(n)
	n.fnChan = make(chan func(), 4096)
	n.ctx, n.cancel = context.WithCancel(o.ctx)

	return n
}

// Name 组件名称
func (n *Node) Name() string {
	return n.opts.name
}

// Init 初始化节点
func (n *Node) Init() {
	if n.opts.id == "" {
		log.Fatal("instance id can not be empty")
	}

	if n.opts.codec == nil {
		log.Fatal("codec component is not injected")
	}

	if n.opts.locator == nil {
		log.Fatal("locator component is not injected")
	}

	if n.opts.registry == nil {
		log.Fatal("registry component is not injected")
	}

	if n.opts.transporter == nil {
		log.Fatal("transporter component is not injected")
	}
}

// Start 启动节点
func (n *Node) Start() {
	n.setState(cluster.Work)

	n.opts.transporter.SetDefaultDiscovery(n.opts.registry)

	n.startRPCServer()

	n.registerServiceInstance()

	n.proxy.watch(n.ctx)

	n.startPromServer()

	n.startCatServer()

	go n.dispatch()

	n.debugPrint()
}

// Destroy 销毁网关服务器
func (n *Node) Destroy() {
	n.deregisterServiceInstance()

	n.stopRPCServer()

	n.stopPromServer()

	n.stopCatServer()

	n.events.close()

	n.router.close()

	close(n.fnChan)

	n.cancel()
}

// Proxy 获取节点代理
func (n *Node) Proxy() *Proxy {
	return n.proxy
}

// 分发处理消息
func (n *Node) dispatch() {
	for {
		select {
		case evt, ok := <-n.events.receive():
			if !ok {
				return
			}
			log.Debugf("node handle receive event [%v], with GID [%v], UID [%v], CID [%v]",
				evt.Event, evt.GID, evt.UID, evt.CID)
			if n.opts.asyncEventHandle {
				go func() {
					xcall.Call(func() {
						n.events.handle(evt)
					})
				}()
			} else {
				xcall.Call(func() {
					n.events.handle(evt)
				})
			}
		case ctx, ok := <-n.router.receive():
			if !ok {
				return
			}
			log.Debugf("node handle receive route [%v] Seq:[%v] with GID [%v], NID [%v], UID [%v], CID [%v]",
				ctx.Request.Message.Route, ctx.Request.Message.Seq, ctx.Request.GID, ctx.Request.NID, ctx.Request.UID, ctx.Request.CID)
			if n.opts.asyncRouterHandle {
				go func() {
					xcall.Call(func() {
						n.router.handle(ctx)
					})
				}()
			} else {
				xcall.Call(func() {
					n.router.handle(ctx)
				})
			}
		case handle, ok := <-n.fnChan:
			if !ok {
				return
			}
			xcall.Call(handle)
		}
	}
}

func (n *Node) startPromServer() {
	n.opts.promServer.Start()
}

func (n *Node) stopPromServer() {
	n.opts.promServer.Destroy()
}

func (n *Node) startCatServer() {
	if n.opts.catServer != nil {
		n.opts.catServer.Start()
	}
}

func (n *Node) stopCatServer() {
	if n.opts.catServer != nil {
		n.opts.catServer.Destroy()
	}
}

// 启动RPC服务器
func (n *Node) startRPCServer() {
	var err error

	n.rpc, err = n.opts.transporter.NewNodeServer(&provider{n})
	if err != nil {
		log.Fatalf("rpc server create failed: %v", err)
	}

	go func() {
		if err = n.rpc.Start(); err != nil {
			log.Fatalf("rpc server start failed: %v", err)
		}
	}()
}

// 停止RPC服务器
func (n *Node) stopRPCServer() {
	if err := n.rpc.Stop(); err != nil {
		log.Errorf("rpc server stop failed: %v", err)
	}
}

// 注册服务实例
func (n *Node) registerServiceInstance() {
	routes := make([]registry.Route, 0, len(n.router.routes))
	for _, entity := range n.router.routes {
		routes = append(routes, registry.Route{
			ID:       entity.route,
			Stateful: entity.stateful,
		})
	}

	events := make([]cluster.Event, 0, len(n.events.events))
	for event := range n.events.events {
		events = append(events, event)
	}

	n.instance = &registry.ServiceInstance{
		ID:        n.opts.id,
		Name:      fmt.Sprintf("%s%s", n.opts.namespace, string(cluster.Node)),
		Kind:      cluster.Node,
		Alias:     n.opts.name,
		State:     n.getState(),
		Routes:    routes,
		Events:    events,
		Endpoint:  n.rpc.Endpoint().String(),
		MetaMap:   n.opts.registry.GetMetaMap(),
		Namespace: n.opts.namespace,
	}
	if n.opts.promServer.Enable() {
		metricsPort, err := strconv.Atoi(n.opts.promServer.GetMetricsPort())
		if err != nil {
			panic(err)
		}
		n.instance.MetricsPort = metricsPort
	}

	ctx, cancel := context.WithTimeout(n.ctx, 10*time.Second)
	err := n.opts.registry.Register(ctx, n.instance)
	cancel()
	if err != nil {
		log.Fatalf("register dispatcher instance failed: %v", err)
	}
}

// 解注册服务实例
func (n *Node) deregisterServiceInstance() {
	ctx, cancel := context.WithTimeout(n.ctx, 10*time.Second)
	err := n.opts.registry.Deregister(ctx, n.instance)
	cancel()
	if err != nil {
		log.Errorf("deregister dispatcher instance failed: %v", err)
	}
}

// 设置节点状态
func (n *Node) setState(state cluster.State) {
	if n.checkState(state) {
		return
	}

	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&n.state)), unsafe.Pointer(&state))

	if n.instance != nil {
		n.instance.State = n.getState()
		for i := 0; i < 3; i++ {
			ctx, cancel := context.WithTimeout(n.ctx, 10*time.Second)
			err := n.opts.registry.Register(ctx, n.instance)
			cancel()
			if err == nil {
				break
			}
		}
	}

	return
}

func (n *Node) addMetaMap(key string, value string) {
	if n.instance != nil {
		n.instance.MetaMap[key] = value
		for i := 0; i < 3; i++ {
			ctx, cancel := context.WithTimeout(n.ctx, 10*time.Second)
			err := n.opts.registry.Register(ctx, n.instance)
			cancel()
			if err == nil {
				break
			}
		}
	}

	return
}

// 获取节点状态
func (n *Node) getState() cluster.State {
	if state := (*cluster.State)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&n.state)))); state == nil {
		return cluster.Shut
	} else {
		return *state
	}
}

// 检测节点状态
func (n *Node) checkState(state cluster.State) bool {
	return n.getState() == state
}

func (n *Node) debugPrint() {
	log.Debugf("node server startup successful")
	log.Debugf("%s server listen on %s", n.rpc.Scheme(), n.rpc.Addr())
}
