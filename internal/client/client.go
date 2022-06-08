package client

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/service/internal/ancillary"
	mw_codec "github.com/go-serv/service/internal/grpc/middleware/codec"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"net"
)

type client struct {
	svcName      protoreflect.FullName
	codec        z.CodecInterface
	codecMwGroup z.CodecMiddlewareGroupInterface
	mwGroup      z.MiddlewareInterface
	endpoint     z.EndpointInterface
	conn         net.Conn
	dialOpts     []grpc.DialOption
	ancillary.MethodMustBeImplemented
}

func (c *client) ServiceName() protoreflect.FullName {
	return c.svcName
}

func (c *client) Codec() z.CodecInterface {
	return c.codec
}

func (s *client) WithCodec(cc z.CodecInterface) {
	s.codec = cc
	s.codecMwGroup = mw_codec.NewCodecMiddlewareGroup(cc)
}

func (s *client) CodecMiddlewareGroup() z.CodecMiddlewareGroupInterface {
	return s.codecMwGroup
}

func (c *client) Endpoint() z.EndpointInterface {
	return c.endpoint
}

func (c *client) WithDialOption(opts grpc.DialOption) {
	c.dialOpts = append(c.dialOpts, opts)
}

func (c *client) DialOptions() []grpc.DialOption {
	return c.dialOpts
}

func (c *client) NewClient(cc grpc.ClientConnInterface) {
	c.MethodMustBeImplemented.Panic()
}

func (c *client) ConnectTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	return nil, nil, nil
}
