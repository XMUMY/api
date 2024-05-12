package v4

var (
	UserNotFoundError  error
	NoCustomTokenError error
	NeedRegisterError  error
)

func init() {
	file_user_v4_error_reason_proto_init()
	UserNotFoundError = ErrorUserNotFound("")
	NoCustomTokenError = ErrorNoCustomToken("")
	NeedRegisterError = ErrorNeedRegister("")
}
