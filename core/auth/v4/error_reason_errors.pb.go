// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v4

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUnknown(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UNKNOWN.String() && e.Code == 500
}

func ErrorUnknown(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_UNKNOWN.String(), fmt.Sprintf(format, args...))
}

func IsInvalidCredential(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_INVALID_CREDENTIAL.String() && e.Code == 400
}

func ErrorInvalidCredential(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_INVALID_CREDENTIAL.String(), fmt.Sprintf(format, args...))
}

func IsInvalidCampusId(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_INVALID_CAMPUS_ID.String() && e.Code == 400
}

func ErrorInvalidCampusId(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_INVALID_CAMPUS_ID.String(), fmt.Sprintf(format, args...))
}

func IsInvalidPassword(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_INVALID_PASSWORD.String() && e.Code == 400
}

func ErrorInvalidPassword(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_INVALID_PASSWORD.String(), fmt.Sprintf(format, args...))
}

func IsWrongPassword(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_WRONG_PASSWORD.String() && e.Code == 403
}

func ErrorWrongPassword(format string, args ...interface{}) *errors.Error {
	return errors.New(403, ErrorReason_WRONG_PASSWORD.String(), fmt.Sprintf(format, args...))
}

func IsLdapServerDown(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_LDAP_SERVER_DOWN.String() && e.Code == 500
}

func ErrorLdapServerDown(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_LDAP_SERVER_DOWN.String(), fmt.Sprintf(format, args...))
}

func IsWrongFirebaseIdToken(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_WRONG_FIREBASE_ID_TOKEN.String() && e.Code == 403
}

func ErrorWrongFirebaseIdToken(format string, args ...interface{}) *errors.Error {
	return errors.New(403, ErrorReason_WRONG_FIREBASE_ID_TOKEN.String(), fmt.Sprintf(format, args...))
}