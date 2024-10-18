package node

import (
	"context"
	cli "github.com/smallnest/rpcx/client"
	"github.com/symsimmy/due/transport"
	"github.com/symsimmy/due/transport/rpcx/internal/code"
	"github.com/symsimmy/due/transport/rpcx/internal/protocol"
)

type Client struct {
	cli *cli.OneClient
}

func NewClient(cli *cli.OneClient) *Client {
	return &Client{cli: cli}
}

// Trigger 触发事件
func (c *Client) Trigger(ctx context.Context, args *transport.TriggerArgs) (miss bool, err error) {
	req := &protocol.TriggerRequest{Event: args.Event, GID: args.GID, CID: args.CID, UID: args.UID}
	reply := &protocol.TriggerReply{}
	err = c.cli.Call(ctx, ServicePath, serviceTriggerMethod, req, reply)
	miss = reply.Code == code.NotFoundSession

	return
}

// Deliver 投递消息
func (c *Client) Deliver(ctx context.Context, args *transport.DeliverArgs) (miss bool, err error) {
	req := &protocol.DeliverRequest{
		GID: args.GID,
		NID: args.NID,
		CID: args.CID,
		UID: args.UID,
		Message: &protocol.Message{
			Seq:        args.Message.Seq,
			Route:      args.Message.Route,
			Buffer:     args.Message.Buffer,
			KcpChannel: protocol.KcpChannel(args.Message.KcpChannel),
		}}
	reply := &protocol.DeliverReply{}
	err = c.cli.Call(ctx, ServicePath, serviceDeliverMethod, req, reply)
	miss = reply.Code == code.NotFoundSession

	return
}
