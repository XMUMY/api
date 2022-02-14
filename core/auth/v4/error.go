package v4

import "github.com/go-kratos/kratos/v2/errors"

var (
	InvalidCredentialError *errors.Error
	InvalidCampusIdError   *errors.Error
	InvalidPasswordError   *errors.Error
	WrongPasswordError     *errors.Error
	LdapServerDown         *errors.Error
	WrongFirebaseIdToken   *errors.Error
)

func init() {
	file_api_auth_v4_error_reason_proto_init()
	InvalidCredentialError = ErrorInvalidCredential("invalid credential")
	InvalidCampusIdError = ErrorInvalidCampusId("invalid campus id")
	InvalidPasswordError = ErrorInvalidPassword("invalid password")
	WrongPasswordError = ErrorWrongPassword("wrong campus id or password")
	LdapServerDown = ErrorLdapServerDown("LDAP service temporary unavailable")
	WrongFirebaseIdToken = ErrorWrongFirebaseIdToken("wrong / revoked firebase id token")
}
