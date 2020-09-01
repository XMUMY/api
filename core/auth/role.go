package auth

import "regexp"

const SvcID = "xmux.core.auth.v4"

var studentIdExp = regexp.MustCompile("^[A-Za-z]{3}[0-9]{7}$")
var stuffIdExp = regexp.MustCompile("^[0-9]{7}$")

// CampusRoleFromUID get campus role from uid.
func CampusRoleFromUID(uid string) CampusRole {
	if matched := studentIdExp.MatchString(uid); matched {
		return CampusRole_student
	} else if matched := stuffIdExp.MatchString(uid); matched {
		return CampusRole_teacher
	} else {
		return CampusRole_none
	}
}