package v4

import (
	"context"
	"encoding/base64"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ExtractAuthorizationFromContext(ctx context.Context) string {
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
	authorization := ExtractAuthorizationFromContext(ctx)

	split := strings.SplitN(authorization, " ", 2)
	if len(split) != 2 || !strings.HasPrefix(strings.ToLower(split[0]), "basic") {
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

// ExtractBearerAuthorizationFromContext and return bearer token.
func ExtractBearerAuthorizationFromContext(ctx context.Context) (token string, err error) {
	authorization := ExtractAuthorizationFromContext(ctx)

	split := strings.SplitN(authorization, " ", 2)
	if len(split) != 2 || !strings.HasPrefix(strings.ToLower(split[0]), "bearer") {
		err = InvalidCredentialError
		return
	}

	token = split[1]
	return
}

type Client interface {
	AuthenticateWithCampusIdPassword(ctx context.Context) (*AuthUserResp, string, error)
	AuthenticateWithFirebaseIdToken(ctx context.Context) (*AuthUserResp, error)
	TryAuthenticate(ctx context.Context) (*AuthUserResp, error)
}

type client struct {
	client AuthInternalClient
}

// NewClient creates a AuthInternalClient wrapper with some helper functions.
func NewClient(cc grpc.ClientConnInterface) Client {
	return &client{
		client: NewAuthInternalClient(cc),
	}
}

// AuthenticateWithCampusIdPassword and return non-nil resp if authentication success.
func (c *client) AuthenticateWithCampusIdPassword(ctx context.Context) (resp *AuthUserResp, password string, err error) {
	var uid string
	uid, password, err = ExtractBasicAuthorizationFromContext(ctx)
	if err != nil {
		return
	}

	resp, err = c.client.AuthUser(
		ctx,
		&AuthUserReq{
			Credential: &AuthUserReq_CampusIdPassword{
				CampusIdPassword: &CampusIdPasswordCredential{
					CampusId: strings.ToLower(uid),
					Password: password,
				},
			},
		},
	)

	return
}

// AuthenticateWithFirebaseIdToken and return non-nil resp if authentication success.
func (c *client) AuthenticateWithFirebaseIdToken(ctx context.Context) (resp *AuthUserResp, err error) {
	var token string
	token, err = ExtractBearerAuthorizationFromContext(ctx)
	if err != nil {
		return
	}

	resp, err = c.client.AuthUser(ctx, &AuthUserReq{
		Credential: &AuthUserReq_FirebaseIdToken{FirebaseIdToken: token},
	})

	return
}

// TryAuthenticate with firebase ID token and campus ID password.
func (c *client) TryAuthenticate(ctx context.Context) (resp *AuthUserResp, err error) {
	token, err := ExtractBearerAuthorizationFromContext(ctx)
	if err == nil && token != "" {
		resp, err = c.client.AuthUser(
			ctx,
			&AuthUserReq{
				Credential: &AuthUserReq_FirebaseIdToken{FirebaseIdToken: token},
			},
		)
		return
	}

	uid, password, err := ExtractBasicAuthorizationFromContext(ctx)
	if err == nil && uid != "" && password != "" {
		resp, err = c.client.AuthUser(
			ctx,
			&AuthUserReq{
				Credential: &AuthUserReq_CampusIdPassword{
					CampusIdPassword: &CampusIdPasswordCredential{
						CampusId: strings.ToLower(uid),
						Password: password,
					},
				},
			},
		)
		return
	}

	return
}
