package net

import (
	"github.com/go-serv/service/internal/client"
	net_cc "github.com/go-serv/service/internal/grpc/codec/net"
	"github.com/go-serv/service/internal/middleware/codec/cipher_msg"
	"github.com/go-serv/service/internal/runtime"
	net_service "github.com/go-serv/service/internal/service/net"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewClient(svcName protoreflect.FullName, e z.EndpointInterface) *netClient {
	c := new(netClient)
	c.insecure = true
	c.ClientInterface = client.NewClient(svcName, e)
	c.svc = net_service.NewNetworkService(svcName)
	// Set client codec
	codec := net_cc.NewOrRegistered(string(svcName))
	c.WithCodec(codec)
	c.WithDialOption(grpc.WithDefaultCallOptions(grpc.ForceCodec(codec)))
	//
	cipher_msg.NetClientInit(c)
	runtime.Runtime().RegisterNetworkClient(c)
	return c
}
