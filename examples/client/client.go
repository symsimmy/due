package main

import (
	"github.com/dobyte/due/log"
	"github.com/dobyte/due/network"
	"github.com/dobyte/due/network/tcp"
	"github.com/dobyte/due/packet"
)

var handlers map[int32]handlerFunc

type handlerFunc func(conn network.Conn, buffer []byte)

func init() {
	handlers = map[int32]handlerFunc{
		1: greetHandler,
	}
}

func main() {
	// 创建客户端
	client := tcp.NewClient()
	// 监听连接
	client.OnConnect(func(conn network.Conn) {
		log.Infof("connection is opened")
	})
	// 监听断开连接
	client.OnDisconnect(func(conn network.Conn) {
		log.Infof("connection is closed")
	})
	// 监听收到消息
	client.OnReceive(func(conn network.Conn, msg []byte, msgType int) {
		message, err := packet.Unpack(msg)
		if err != nil {
			log.Errorf("unpack message failed: %v", err)
			return
		}

		handler, ok := handlers[message.Route]
		if !ok {
			log.Errorf("the route handler is not registered, route:%v", message.Route)
			return
		}
		handler(conn, message.Buffer)
	})

	conn, err := client.Dial()
	if err != nil {
		log.Fatalf("dial failed: %v", err)
	}

	if err = push(conn, 1, []byte("hello due~~")); err != nil {
		log.Errorf("push message failed: %v", err)
	}

	select {}
}

func greetHandler(conn network.Conn, buffer []byte) {
	log.Infof("received message from server: %s", string(buffer))
}

func push(conn network.Conn, route int32, buffer []byte) error {
	msg, err := packet.Pack(&packet.Message{
		Seq:    1,
		Route:  route,
		Buffer: buffer,
	})
	if err != nil {
		return err
	}

	return conn.Push(msg)
}
