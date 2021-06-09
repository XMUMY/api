package v4

import (
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/grpc/codes"
)

var (
	UserNotFoundError = errors.NotFound(ErrorReason_name[int32(ErrorReason_USER_NOT_FOUND)],
		"user not found")

	NoCustomTokenError = errors.ServiceUnavailable(ErrorReason_name[int32(ErrorReason_NO_CUSTOM_TOKEN)],
		"login successful but failed to get firebase token")

	NeedRegisterError = errors.New(int(codes.FailedPrecondition), ErrorReason_name[int32(ErrorReason_NEED_REGISTER)],
		"need register before login") // TODO: GRPC status.
)
