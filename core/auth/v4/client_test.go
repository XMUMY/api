package v4

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"testing"

	"github.com/go-kratos/kratos/v2/transport/grpc"
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

		token := ExtractAuthorizationFromContext(ctx)
		Convey("correct authorization extracted", func() {
			So(token, ShouldEqual, "bearer SOME_TOKEN")
		})
	})
}

func TestExtractBasicAuthorizationFromContext(t *testing.T) {
	Convey("extract basic authorization from context", t, func() {
		encoded := base64.StdEncoding.EncodeToString(
			[]byte(fmt.Sprintf("%s:%s",
				"SOME_USERNAME",
				"SOME_PASSWORD",
			)),
		)
		ctx := metadata.NewIncomingContext(
			context.TODO(),
			metadata.New(map[string]string{
				"authorization": "basic " + encoded,
			}),
		)

		uid, password, err := ExtractBasicAuthorizationFromContext(ctx)
		Convey("correct basic authorization extracted", func() {
			So(uid, ShouldEqual, "SOME_USERNAME")
			So(password, ShouldEqual, "SOME_PASSWORD")
			So(err, ShouldBeNil)
		})
	})
}

func TestExtractBearerAuthorizationFromContext(t *testing.T) {
	Convey("extract bearer authorization from context", t, func() {
		ctx := metadata.NewIncomingContext(
			context.TODO(),
			metadata.New(map[string]string{
				"authorization": "bearer SOME.TOKEN",
			}),
		)

		token, err := ExtractBearerAuthorizationFromContext(ctx)
		Convey("correct bearer authorization extracted", func() {
			So(token, ShouldEqual, "SOME.TOKEN")
			So(err, ShouldBeNil)
		})
	})
}

func TestAuthenticateWithCampusIdPassword(t *testing.T) {
	SkipConvey("authenticate with campus id and password", t, func() {
		encoded := base64.StdEncoding.EncodeToString(
			[]byte(fmt.Sprintf(
				"%s:%s",
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

		conn, err := grpc.DialInsecure(context.TODO(), grpc.WithEndpoint("127.0.0.1:9000"))
		if err != nil {
			panic(err)
		}
		client := NewClient(conn)

		Convey("authenticated successfully", func() {
			resp, password, err := client.AuthenticateWithCampusIdPassword(ctx)
			So(resp, ShouldNotBeNil)
			So(password, ShouldEqual, os.Getenv("TEST_CAMPUS_ID_PASSWORD"))
			So(err, ShouldBeNil)
		})
	})
}
