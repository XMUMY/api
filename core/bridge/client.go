package api

import (
	"context"

	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"google.golang.org/grpc"
)

//go:generate kratos tool protoc

// NewClient grpc client
func NewClient(dsn string, cfg *warden.ClientConfig, opts ...grpc.DialOption) (BridgeClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return NewBridgeClient(cc), nil
}
