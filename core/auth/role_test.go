package auth

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCampusRoleFromUID(t *testing.T) {
	Convey("parse campus role from uid", t, func() {
		Convey("none", func() {
			role := CampusRoleFromUID("TEST989898")
			So(role, ShouldEqual, CampusRole_none)
		})

		Convey("student", func() {
			role := CampusRoleFromUID("TST9898989")
			So(role, ShouldEqual, CampusRole_student)
		})

		Convey("teacher", func() {
			role := CampusRoleFromUID("9898989")
			So(role, ShouldEqual, CampusRole_teacher)
		})
	})
}
