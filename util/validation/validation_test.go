package validation

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEmailValidation(t *testing.T) {

	emailValid := "test@gongan.github.io"
	emailInvalid := "test#gongan.github.io"

	Convey("Check that the email address is valid or not", t, func() {

		Convey("when the valid email checked successfully", func() {
			So(EmailValidation(emailValid), ShouldBeTrue)
		})

		Convey("when the invalid email checked successfully", func() {
			So(EmailValidation(emailInvalid), ShouldBeFalse)
		})
	})
}
