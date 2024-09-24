package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	//GateServerReceiveClientMessageCountCounter 收到 client 消息数
	GateServerReceiveClientMessageCountCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: GateServerReceiveClientMessageCount,
		},
		[]string{"B_server_ip", "B_server_port", "B_route_id"},
	)

	//GateServerReceiveClientMessageBytesGauge 收到 client 消息大小
	GateServerReceiveClientMessageBytesGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: GateServerReceiveClientMessageBytes,
		},
		[]string{"B_server_ip", "B_server_port", "B_route_id"},
	)

	//GateServerSendToClientMessageCountCounter 回给 client 消息数
	GateServerSendToClientMessageCountCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: GateServerSendToClientMessageCount,
		},
		[]string{"B_server_ip", "B_server_port", "B_route_id"},
	)

	//GateServerSendToClientMessageBytesGauge 回给 client 消息大小
	GateServerSendToClientMessageBytesGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: GateServerSendToClientMessageBytes,
		},
		[]string{"B_server_ip", "B_server_port", "B_route_id"},
	)

	//GateServerSendToServerMessageCountCounter 转发内网服务消息数
	// todo 无法获取 nodeId
	GateServerSendToServerMessageCountCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: GateServerSendToServerMessageCount,
		},
		//[]string{"B_routeId", "B_node_id", "B_error_code"},
		[]string{"B_routeId", "B_error_code"},
	)

	//GateServerReceiveFromServerMessageCountCounter 收到内网服务消息数
	// todo nodeId
	GateServerReceiveFromServerMessageCountCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: GateServerReceiveFromServerMessageCount,
		},
		[]string{"B_routeId", "B_node_id"},
	)

	//GateServerTotalOnlinePlayerGauge 在线玩家数量
	// todo
	GateServerTotalOnlinePlayerGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: GateServerTotalOnlinePlayer,
		},
		[]string{"B_gate_instance_id"},
	)

	//GateServerClientReconnectCountCounter 断线重连次数
	// todo
	GateServerClientReconnectCountCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: GateServerClientReconnectCount,
		},
		[]string{},
	)
)
