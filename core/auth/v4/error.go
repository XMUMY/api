package v4

import "github.com/go-kratos/kratos/v2/errors"

const domain = "api.xmux.xdea.io"

var (
	InvalidCredentialError = errors.Unauthorized(domain, ErrorReason_name[int32(ErrorReason_INVALID_CREDENTIAL)],
		"invalid credential")

	InvalidCampusIdError = errors.BadRequest(domain, ErrorReason_name[int32(ErrorReason_INVALID_CAMPUS_ID)],
		"invalid campus id")

	InvalidPasswordError = errors.BadRequest(domain, ErrorReason_name[int32(ErrorReason_INVALID_PASSWORD)],
		"invalid password")

	WRONG_PASSWORD_ERROR = errors.Forbidden(domain, ErrorReason_name[int32(ErrorReason_WRONG_PASSWORD)],
		"wrong campus id or password")

	LDAP_SERVER_DOWN = errors.ServiceUnavailable(domain, ErrorReason_name[int32(ErrorReason_LDAP_SERVER_DOWN)],
		"LDAP service temporary unavailable")

	WRONG_FIREBASE_ID_TOKEN = errors.Forbidden(domain,
		ErrorReason_name[int32(ErrorReason_WRONG_FIREBASE_ID_TOKEN)],
		"wrong / revoked firebase id token")
)
