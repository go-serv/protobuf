package net

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc"
)

func (c *netClient) ConnectTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	init := func(task job.TaskInterface) {
		netEndpoint := c.Endpoint().(z.NetEndpointInterface)
		transCreds := netEndpoint.TransportCredentials()
		c.WithDialOption(grpc.WithTransportCredentials(transCreds))
	}
	run := func(task job.TaskInterface) {
		v := j.GetValue()
		opts := append(c.DialOptions(), grpc.WithChainUnaryInterceptor(c.Middleware().UnaryClientInterceptor()))
		conn, err := grpc.Dial(c.Endpoint().Address(), opts...)
		task.Assert(err)
		v.(z.NetworkClientInterface).OnConnect(conn)
		task.Done()
	}
	return init, run, nil
}