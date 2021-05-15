package v4

import "github.com/go-kratos/kratos/v2/errors"

const domain = "api.xmux.xdea.io"

var (
	ACServiceUnavailableError = errors.ServiceUnavailable(domain, ErrorReason_name[int32(ErrorReason_AC_SERVICE_UNAVAILABLE)],
		"AC service temporary unavailable")

	ACWrongPasswordError = errors.Forbidden(domain, ErrorReason_name[int32(ErrorReason_AC_WRONG_PASSWORD)],
		"wrong password")
)
