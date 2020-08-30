package auth

import (
	"context"
	"encoding/base64"
	"strings"

	"github.com/micro/go-micro/v2/service"
	"google.golang.org/grpc/metadata"
)

var svc AuthService

// InitAuthService from current service client.
func InitAuthService(from service.Service) {
	svc = NewAuthService(SvcName, from.Client())
}

func extractAuthorizationFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	authorization, ok := md["authorization"]
	if !ok || len(authorization) == 0 {
		return ""
	}

	return authorization[0]
}

// AuthenticateWithCampusIdPassword and return non-nil if authentication success.
func AuthenticateWithCampusIdPassword(ctx context.Context) *AuthUserResp {
	authorization := extractAuthorizationFromContext(ctx)
	if authorization == "" {
		return nil
	}
	return authenticateWithCampusIdPassword(ctx, authorization)
}

func authenticateWithCampusIdPassword(ctx context.Context, authorization string) *AuthUserResp {
	split := strings.SplitN(authorization, " ", 2)
	if !strings.HasPrefix(strings.ToLower(split[0]), "basic") {
		return nil
	}

	decoded, err := base64.StdEncoding.DecodeString(split[1])
	if err != nil {
		return nil
	}

	split = strings.SplitN(string(decoded), ":", 2)
	if len(split) != 2 {
		return nil
	}

	resp, err := svc.AuthUser(ctx, &AuthUserReq{
		Credential: &AuthUserReq_CampusIdPassword{
			CampusIdPassword: &CampusIdPasswordCredential{
				CampusId: strings.ToLower(split[0]),
				Password: split[1],
			},
		},
	})

	return resp
}
