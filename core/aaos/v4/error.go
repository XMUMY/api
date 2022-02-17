package v4

import "github.com/go-kratos/kratos/v2/errors"

var (
	ACServiceUnavailableError *errors.Error
	ACWrongPasswordError      *errors.Error
)

func init() {
	file_api_aaos_v4_error_reason_proto_init()
	ACServiceUnavailableError = ErrorAcServiceUnavailable("AC service temporary unavailable")
	ACWrongPasswordError = ErrorAcWrongPassword("wrong password")
}
