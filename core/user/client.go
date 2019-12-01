package api

import (
	"context"
	"os"

	"git.xdea.xyz/micros/lib/net/metadata"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

//go:generate kratos tool protoc

// NewClient create new UserClient.
func NewClient(dsn string, cfg *warden.ClientConfig, opts ...grpc.DialOption) (UserClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return NewUserClient(cc), nil
}

var userAPIClient UserClient

// ensureUserClient existed. Or will create default auth client by env USER_ADDR.
func ensureUserClient(client UserClient) {
	if client != nil {
		userAPIClient = client
	} else if userAPIClient == nil {
		authAddr := os.Getenv("USER_ADDR")
		if authAddr == "" {
			panic("user service address not found")
		}
		var err error
		userAPIClient, err = NewClient(authAddr, &warden.ClientConfig{})
		if err != nil {
			log.Error("failed to connect user service")
			return
		}
	}
}

// AuthFirebaseJWT will create a blademaster middleware to verify firebase JWT token if exist.
//
// UserClient will be created by env AUTH_ADDR if not specified.
func AuthFirebaseJWT(client UserClient) blademaster.HandlerFunc {
	ensureUserClient(client)

	return func(ctx *blademaster.Context) {
		_, err := metadata.BearerFromContext(ctx)
		if err == nil {
			c := metadata.TransferAuthorizationToGRPC(ctx)
			authedUser, err := userAPIClient.AuthUserWithFirebaseJWT(c, &empty.Empty{})
			if err == nil {
				ctx.Set("authedUser", authedUser)
			}
		}
		ctx.Next()
	}
}

// ForceAuthFirebaseJWT will create a blademaster middleware to verify firebase JWT token
// and abort with 403/401 if failed to authenticate / not exist.
//
// UserClient will be created by env AUTH_ADDR if not specified.
func ForceAuthFirebaseJWT(client UserClient) blademaster.HandlerFunc {
	ensureUserClient(client)

	return func(ctx *blademaster.Context) {
		_, err := metadata.BearerFromContext(ctx)
		if err != nil {
			ctx.AbortWithStatus(401)
			return
		}

		c := metadata.TransferAuthorizationToGRPC(ctx)
		authedUser, err := userAPIClient.AuthUserWithFirebaseJWT(c, &empty.Empty{})
		if err != nil {
			ctx.AbortWithStatus(403)
			return
		}

		ctx.Set("authedUser", authedUser)
		ctx.Next()
	}
}

// EnsureAuthenticated try to get the authenticated user from context. Return nil if not found.
//
// UserClient will be created by env USER_ADDR if not specified.
func EnsureAuthenticated(ctx context.Context, client UserClient) (authedUser *AuthedUser) {
	ensureUserClient(client)

	if bmContext, ok := ctx.(blademaster.Context); ok {
		if contextUser, ok := bmContext.Get("authedUser"); ok {
			return contextUser.(*AuthedUser)
		}
	}

	uid, password, err := metadata.BasicFromContext(ctx)
	if err == nil && uid != "" && password != "" {
		ctx = metadata.TransferAuthorizationToGRPC(ctx)
		authedUser, err = userAPIClient.AuthUser(ctx, &empty.Empty{})
		if err == nil {
			return
		}
	}

	token, err := metadata.BearerFromContext(ctx)
	if err == nil && token != "" {
		ctx = metadata.TransferAuthorizationToGRPC(ctx)
		authedUser, err = userAPIClient.AuthUserWithFirebaseJWT(ctx, &empty.Empty{})
		if err == nil {
			return
		}
	}

	return nil
}
