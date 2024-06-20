package node

import (
	"context"
	"github.com/symsimmy/due/config"
	"github.com/symsimmy/due/crypto"
	"github.com/symsimmy/due/encoding"
	"github.com/symsimmy/due/locate"
	"github.com/symsimmy/due/metrics/cat"
	"github.com/symsimmy/due/metrics/prometheus"
	"github.com/symsimmy/due/registry"
	"github.com/symsimmy/due/transport"
	"github.com/symsimmy/due/utils/xuuid"
	"time"
)

const (
	defaultName                    = "node"          // 默认节点名称
	defaultCodec                   = "proto"         // 默认编解码器名称
	defaultTimeout                 = 3 * time.Second // 默认超时时间
	defaultEnableAsyncEventHandle  = true            // event事件异步执行
	defaultEnableAsyncRouterHandle = true            // route事件异步执行
	defaultNamespace               = ""
)

const (
	defaultIDKey                = "config.cluster.node.id"
	defaultNameKey              = "config.cluster.node.name"
	defaultCodecKey             = "config.cluster.node.codec"
	defaultTimeoutKey           = "config.cluster.node.timeout"
	defaultEncryptorKey         = "config.cluster.node.encryptor"
	defaultDecryptorKey         = "config.cluster.node.decryptor"
	defaultAsyncEventHandleKey  = "config.cluster.node.async_event_handle"
	defaultAsyncRouterHandleKey = "config.cluster.node.async_router_handle"
	defaultNamespaceKey         = "config.cluster.node.namespace"
)

type Option func(o *options)

type options struct {
	id                string                // 实例ID
	name              string                // 实例名称
	ctx               context.Context       // 上下文
	codec             encoding.Codec        // 编解码器
	timeout           time.Duration         // RPC调用超时时间
	locator           locate.Locator        // 用户定位器
	registry          registry.Registry     // 服务注册器
	transporter       transport.Transporter // 消息传输器
	promServer        prometheus.PromServer // 埋点采集服务器
	catServer         *cat.Server           // cat服务器
	encryptor         crypto.Encryptor      // 消息加密器
	decryptor         crypto.Decryptor      // 消息解密器
	asyncEventHandle  bool                  // 异步处理连接事件
	asyncRouterHandle bool                  // 异步处理消息事件
	namespace         string
}

func defaultOptions() *options {
	opts := &options{
		ctx:       context.Background(),
		name:      defaultName,
		codec:     encoding.Invoke(defaultCodec),
		timeout:   defaultTimeout,
		namespace: config.Get(defaultNamespaceKey, defaultNamespace).String(),
	}

	if id := config.Get(defaultIDKey).String(); id != "" {
		opts.id = id
	} else if id, err := xuuid.UUID(); err == nil {
		opts.id = id
	}

	if name := config.Get(defaultNameKey).String(); name != "" {
		opts.name = name
	}

	if codec := config.Get(defaultCodecKey).String(); codec != "" {
		opts.codec = encoding.Invoke(codec)
	}

	if timeout := config.Get(defaultTimeoutKey).Int64(); timeout > 0 {
		opts.timeout = time.Duration(timeout) * time.Second
	}

	if encryptor := config.Get(defaultEncryptorKey).String(); encryptor != "" {
		opts.encryptor = crypto.InvokeEncryptor(encryptor)
	}

	if decryptor := config.Get(defaultDecryptorKey).String(); decryptor != "" {
		opts.decryptor = crypto.InvokeDecryptor(decryptor)
	}

	opts.asyncEventHandle = config.Get(defaultAsyncEventHandleKey, true).Bool()
	opts.asyncRouterHandle = config.Get(defaultAsyncRouterHandleKey, true).Bool()

	return opts
}

// WithID 设置实例ID
func WithID(id string) Option {
	return func(o *options) { o.id = id }
}

// WithName 设置实例名称
func WithName(name string) Option {
	return func(o *options) { o.name = name }
}

// WithCodec 设置编解码器
func WithCodec(codec encoding.Codec) Option {
	return func(o *options) { o.codec = codec }
}

// WithContext 设置上下文
func WithContext(ctx context.Context) Option {
	return func(o *options) { o.ctx = ctx }
}

// WithTimeout 设置RPC调用超时时间
func WithTimeout(timeout time.Duration) Option {
	return func(o *options) { o.timeout = timeout }
}

// WithLocator 设置定位器
func WithLocator(locator locate.Locator) Option {
	return func(o *options) { o.locator = locator }
}

// WithRegistry 设置服务注册器
func WithRegistry(r registry.Registry) Option {
	return func(o *options) { o.registry = r }
}

// WithTransporter 设置消息传输器
func WithTransporter(transporter transport.Transporter) Option {
	return func(o *options) { o.transporter = transporter }
}

// WithPromServer 设置prom server
func WithPromServer(promServer *prometheus.PromServer) Option {
	return func(o *options) { o.promServer = *promServer }
}

// WithCatServer 设置cat server
func WithCatServer(catServer *cat.Server) Option {
	return func(o *options) { o.catServer = catServer }
}

// WithEncryptor 设置消息加密器
func WithEncryptor(encryptor crypto.Encryptor) Option {
	return func(o *options) { o.encryptor = encryptor }
}

// WithDecryptor 设置消息解密器
func WithDecryptor(decryptor crypto.Decryptor) Option {
	return func(o *options) { o.decryptor = decryptor }
}
