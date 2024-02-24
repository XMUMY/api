package v4

var (
	InvalidCredentialError    error
	InvalidCampusIdError      error
	InvalidPasswordError      error
	WrongPasswordError        error
	WrongFirebaseIdTokenError error
	LdapServerDownError       error
)

func init() {
	file_auth_v4_error_reason_proto_init()
	InvalidCredentialError = ErrorInvalidCredential("")
	InvalidCampusIdError = ErrorInvalidCampusId("")
	InvalidPasswordError = ErrorInvalidPassword("")
	WrongPasswordError = ErrorWrongPassword("")
	WrongFirebaseIdTokenError = ErrorWrongFirebaseIdToken("")
	LdapServerDownError = ErrorLdapServerDown("")
}
