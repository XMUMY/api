package auth

import (
	"context"
	"encoding/base64"
	"strings"

	"github.com/micro/go-micro/v2"
	"google.golang.org/grpc/metadata"
)

var svc AuthService

// InitAuthService from current service client.
func InitAuthService(from micro.Service) {
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

// ExtractBasicAuthorizationFromContext and return basic credentials.
func ExtractBasicAuthorizationFromContext(ctx context.Context) (username string, password string, err error) {
	authorization := extractAuthorizationFromContext(ctx)

	split := strings.SplitN(authorization, " ", 2)
	if !strings.HasPrefix(strings.ToLower(split[0]), "basic") {
		err = InvalidCredentialError
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(split[1])
	if err != nil {
		err = InvalidCredentialError
		return
	}

	split = strings.SplitN(string(decoded), ":", 2)
	if len(split) != 2 {
		err = InvalidCredentialError
		return
	}

	username = split[0]
	password = split[1]
	return
}

// AuthenticateWithCampusIdPassword and return non-nil resp if authentication success.
func AuthenticateWithCampusIdPassword(ctx context.Context) (resp *AuthUserResp, password string, err error) {
	var uid string
	uid, password, err = ExtractBasicAuthorizationFromContext(ctx)
	if err != nil {
		return
	}

	resp, err = svc.AuthUser(ctx, &AuthUserReq{
		Credential: &AuthUserReq_CampusIdPassword{
			CampusIdPassword: &CampusIdPasswordCredential{
				CampusId: strings.ToLower(uid),
				Password: password,
			},
		},
	})

	return
}
