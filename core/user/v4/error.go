package v4

import (
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/grpc/codes"
)

const domain = "api.xmux.xdea.io"

var (
	UserNotFoundError = errors.NotFound(domain, ErrorReason_name[int32(ErrorReason_USER_NOT_FOUND)],
		"user not found")

	NoCustomTokenError = errors.ServiceUnavailable(domain, ErrorReason_name[int32(ErrorReason_NO_CUSTOM_TOKEN)],
		"login successful but failed to get firebase token")

	NeedRegisterError = errors.New(codes.FailedPrecondition, domain, ErrorReason_name[int32(ErrorReason_NEED_REGISTER)],
		"need register before login")
)
