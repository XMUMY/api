package v4

import "github.com/go-kratos/kratos/v2/errors"

var (
	InvalidCredentialError = errors.Unauthorized(ErrorReason_name[int32(ErrorReason_INVALID_CREDENTIAL)],
		"invalid credential")

	InvalidCampusIdError = errors.BadRequest(ErrorReason_name[int32(ErrorReason_INVALID_CAMPUS_ID)],
		"invalid campus id")

	InvalidPasswordError = errors.BadRequest(ErrorReason_name[int32(ErrorReason_INVALID_PASSWORD)],
		"invalid password")

	WrongPasswordError = errors.Forbidden(ErrorReason_name[int32(ErrorReason_WRONG_PASSWORD)],
		"wrong campus id or password")

	LdapServerDown = errors.ServiceUnavailable(ErrorReason_name[int32(ErrorReason_LDAP_SERVER_DOWN)],
		"LDAP service temporary unavailable")

	WrongFirebaseIdToken = errors.Forbidden(ErrorReason_name[int32(ErrorReason_WRONG_FIREBASE_ID_TOKEN)],
		"wrong / revoked firebase id token")
)
