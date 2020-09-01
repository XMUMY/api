package auth

import (
	"net/http"

	"github.com/micro/go-micro/v2/errors"
)

var (
	InvalidCredentialError = errors.Unauthorized(SvcID, "invalid credential")

	InvalidCampusIdError = errors.BadRequest(SvcID, "invalid campus id")

	InvalidPasswordError = errors.BadRequest(SvcID, "invalid password")

	LDAPPasswordError = &errors.Error{
		Id:     SvcID,
		Code:   http.StatusForbidden,
		Detail: "wrong campus id or password",
		Status: http.StatusText(http.StatusForbidden),
	}

	LDAPServerError = &errors.Error{
		Id:     SvcID,
		Code:   http.StatusServiceUnavailable,
		Detail: "LDAP service temporary unavailable",
		Status: http.StatusText(http.StatusServiceUnavailable),
	}

	FirebaseIdTokenError = &errors.Error{
		Id:     SvcID,
		Code:   http.StatusForbidden,
		Detail: "wrong / revoked id token",
		Status: http.StatusText(http.StatusForbidden),
	}
)