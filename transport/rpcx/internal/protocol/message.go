package protocol

type Message struct {
	Seq        int32  // 序列号
	Route      int32  // 路由
	Buffer     []byte // 消息内容
	KcpChannel KcpChannel
}

type KcpChannel int32

const (
	Reliable   KcpChannel = iota // Reliable = 0
	Unreliable                   // Unreliable = 1
)
