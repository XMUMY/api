package v4

var (
	AcServiceUnavailableError error
	AcWrongPasswordError      error
	AcAccountBlockedError     error
)

func init() {
	file_aaos_v4_error_reason_proto_init()
	AcServiceUnavailableError = ErrorAcServiceUnavailable("")
	AcWrongPasswordError = ErrorAcWrongPassword("")
	AcAccountBlockedError = ErrorAcAccountBlocked("")
}
