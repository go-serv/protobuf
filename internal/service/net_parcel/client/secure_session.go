package client

import (
	"context"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"google.golang.org/grpc"
)

func (c *client) SecureSession(ctx context.Context, in *proto.Session_Request, opts ...grpc.CallOption) (*proto.Session_Response, error) {
	return c.stubs.SecureSession(ctx, in)
}