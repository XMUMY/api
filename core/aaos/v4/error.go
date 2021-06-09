package v4

import "github.com/go-kratos/kratos/v2/errors"

var (
	ACServiceUnavailableError = errors.ServiceUnavailable(ErrorReason_name[int32(ErrorReason_AC_SERVICE_UNAVAILABLE)],
		"AC service temporary unavailable")

	ACWrongPasswordError = errors.Forbidden(ErrorReason_name[int32(ErrorReason_AC_WRONG_PASSWORD)],
		"wrong password")
)
