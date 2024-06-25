package consul

import (
	"context"
	"fmt"
	"github.com/symsimmy/due/cluster"
	"github.com/symsimmy/due/encoding/json"
	"github.com/symsimmy/due/log"
	"net"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/hashicorp/consul/api"

	"github.com/symsimmy/due/registry"
)

var _ registry.Registry = &Registry{}

type Registry struct {
	err        error
	ctx        context.Context
	cancel     context.CancelFunc
	opts       *options
	watchers   sync.Map
	registrars sync.Map
}

func NewRegistry(opts ...Option) *Registry {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	r := &Registry{}
	r.opts = o
	r.ctx, r.cancel = context.WithCancel(o.ctx)

	if o.client == nil {
		config := api.DefaultConfig()
		if o.addr != "" {
			config.Address = o.addr
		}

		o.client, r.err = api.NewClient(config)
	}

	return r
}

func (r *Registry) GetMetaMap() map[string]string {
	return r.opts.metaMap
}

// Register 注册服务实例
func (r *Registry) Register(ctx context.Context, ins *registry.ServiceInstance) error {
	if r.err != nil {
		return r.err
	}

	v, ok := r.registrars.Load(ins.ID)
	if ok {
		return v.(*registrar).register(ctx, ins)
	}

	reg := newRegistrar(r)

	if err := reg.register(ctx, ins); err != nil {
		return err
	}

	r.registrars.Store(ins.ID, reg)

	return nil
}

// Deregister 解注册服务实例
func (r *Registry) Deregister(ctx context.Context, ins *registry.ServiceInstance) error {
	v, ok := r.registrars.Load(ins.ID)
	if ok {
		return v.(*registrar).deregister(ctx, ins)
	}

	return r.opts.client.Agent().ServiceDeregister(ins.ID)
}

// Services 获取服务实例列表
func (r *Registry) Services(ctx context.Context, namespace string, kind cluster.Kind, alias string) ([]*registry.ServiceInstance, error) {
	serviceName := r.getServicesName(namespace, kind, alias)
	if r.err != nil {
		return nil, r.err
	}

	v, ok := r.watchers.Load(serviceName)
	if ok {
		return v.(*watcherMgr).services(), nil
	} else {
		services, _, err := r.services(ctx, serviceName, 0, true)
		return services, err
	}
}

// Watch 监听服务
func (r *Registry) Watch(ctx context.Context, namespace string, kind cluster.Kind, alias string) (registry.Watcher, error) {
	serviceName := r.getServicesName(namespace, kind, alias)
	if r.err != nil {
		return nil, r.err
	}

	v, ok := r.watchers.Load(serviceName)
	if ok {
		return v.(*watcherMgr).fork(), nil
	}

	w, err := newWatcherMgr(r, ctx, serviceName)
	if err != nil {
		return nil, err
	}
	r.watchers.Store(serviceName, w)

	return w.fork(), nil
}

func (r *Registry) getServicesName(namespace string, kind cluster.Kind, alias string) string {
	if kind == cluster.Node {
		return fmt.Sprintf("%s%s.%s", namespace, string(kind), alias)
	} else {
		return fmt.Sprintf("%s%s", namespace, string(kind))
	}
}

// 获取服务实体列表
func (r *Registry) services(ctx context.Context, serviceName string, waitIndex uint64, passingOnly bool) ([]*registry.ServiceInstance, uint64, error) {
	opts := &api.QueryOptions{
		WaitIndex: waitIndex,
		WaitTime:  60 * time.Second,
	}
	opts.WithContext(ctx)

	entries, meta, err := r.opts.client.Health().Service(serviceName, "", passingOnly, opts)
	if err != nil {
		return nil, 0, err
	}

	services := make([]*registry.ServiceInstance, 0, len(entries))
	for _, entry := range entries {
		ins := &registry.ServiceInstance{
			ID:     entry.Service.ID,
			Name:   entry.Service.Service,
			Routes: make([]registry.Route, 0, len(entry.Service.Meta)),
		}

		for scheme, addr := range entry.Service.TaggedAddresses {
			if scheme == "lan_ipv4" || scheme == "wan_ipv4" || scheme == "lan_ipv6" || scheme == "wan_ipv6" {
				continue
			}
			ins.Endpoint = (&url.URL{
				Scheme: scheme,
				Host:   net.JoinHostPort(addr.Address, strconv.Itoa(addr.Port)),
			}).String()
		}
		if ins.Endpoint == "" {
			continue
		}

		for k, v := range entry.Service.Meta {
			switch k {
			case metaFieldKind:
				ins.Kind = cluster.Kind(v)
			case metaFieldAlias:
				ins.Alias = v
			case metaFieldState:
				ins.State = cluster.State(v)
			default:
				if ins.MetaMap == nil {
					ins.MetaMap = make(map[string]string)
				}
				ins.MetaMap[k] = v
			}
		}

		for _, v := range entry.Service.Tags {
			event, err := strconv.Atoi(v)
			if err != nil {
				continue
			}

			ins.Events = append(ins.Events, cluster.Event(event))
		}

		// get route in key/value
		res, _, err := r.opts.client.KV().Get(fmt.Sprintf(routeKvFormat, ins.Alias, ins.ID), nil)
		if res == nil || err != nil {
			log.Warnf("get [%v]:[%v] route failed, caused by [%+v] or instance [%v] is not found", ins.Alias, ins.ID, err, ins.ID)
			continue
		}

		var m map[int32]bool
		err = json.Unmarshal(res.Value, &m)
		if err != nil {
			log.Warnf("unmarshal [%v]:[%v] route failed, caused by [%+v]", ins.Alias, ins.ID, err)
			continue
		}

		for key, value := range m {
			route := key
			stateful := value

			ins.Routes = append(ins.Routes, registry.Route{
				ID:       route,
				Stateful: stateful,
			})
		}

		services = append(services, ins)
	}

	return services, meta.LastIndex, nil
}
