package api

import (
	"context"
	"os"

	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"google.golang.org/grpc"
)

//go:generate kratos tool protoc

// NewClient create new BridgeClient.
func NewClient(dsn string, cfg *warden.ClientConfig, opts ...grpc.DialOption) (BridgeClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return NewBridgeClient(cc), nil
}

func NewDefaultClient() (BridgeClient, error) {
	return NewClient(os.Getenv("BRIDGE_ADDR"), &warden.ClientConfig{})
}
