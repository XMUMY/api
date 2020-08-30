package auth

import (
	"net/http"

	"github.com/micro/go-micro/v2/errors"
)

var (
	InvalidCredentialError = errors.Unauthorized(SvcName, "invalid credential")

	InvalidCampusIdError = errors.BadRequest(SvcName, "invalid campus id")

	InvalidPasswordError = errors.BadRequest(SvcName, "invalid password")

	LDAPPasswordError = &errors.Error{
		Id:     SvcName,
		Code:   http.StatusForbidden,
		Detail: "wrong campus id or password",
		Status: http.StatusText(http.StatusForbidden),
	}

	LDAPServerError = &errors.Error{
		Id:     SvcName,
		Code:   http.StatusServiceUnavailable,
		Detail: "LDAP service temporary unavailable",
		Status: http.StatusText(http.StatusServiceUnavailable),
	}

	FirebaseIdTokenError = &errors.Error{
		Id:     SvcName,
		Code:   http.StatusForbidden,
		Detail: "wrong / revoked id token",
		Status: http.StatusText(http.StatusForbidden),
	}
)