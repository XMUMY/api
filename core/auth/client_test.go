package auth

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc/metadata"
)

func TestExtractAuthorizationFromContext(t *testing.T) {
	Convey("extract authorization from context", t, func() {
		ctx := metadata.NewIncomingContext(
			context.TODO(),
			metadata.New(map[string]string{
				"authorization": "bearer SOME_TOKEN",
			}),
		)

		token := extractAuthorizationFromContext(ctx)
		Convey("correct authorization extracted", func() {
			So(token, ShouldEqual, "bearer SOME_TOKEN")
		})
	})
}

func TestAuthenticateWithCampusIdPassword(t *testing.T) {
	SkipConvey("authenticate with campus id and password", t, func() {
		encoded := base64.StdEncoding.EncodeToString(
			[]byte(fmt.Sprintf("%s:%s",
				os.Getenv("TEST_CAMPUS_ID"),
				os.Getenv("TEST_CAMPUS_ID_PASSWORD"),
			)),
		)
		ctx := metadata.NewIncomingContext(
			context.TODO(),
			metadata.New(map[string]string{
				"authorization": "basic " + encoded,
			}),
		)

		resp := AuthenticateWithCampusIdPassword(ctx)
		Convey("authenticated successfully", func() {
			So(resp, ShouldNotBeNil)
		})
	})
}