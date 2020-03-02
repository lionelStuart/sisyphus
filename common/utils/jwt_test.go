package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGenToken(t *testing.T) {

	Convey("setup", t, func() {

		t.Log("pass")

		Convey("test gen token", func() {
			claims := Claims{
				Username: "admin",
				Password: "12345",
			}

			jwt, err := GenToken(claims.Username, claims.Password)
			So(err, ShouldBeNil)
			So(jwt, ShouldNotBeEmpty)
			t.Log("token: ", jwt)
		})

		Convey("test parse token", func() {
			var jwtStr = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjIxMjMyZjI5N2E1N2E1YTc0Mzg5NGEwZTRhODAxZmMzIiwicGFzc3dvcmQiOiI4MjdjY2IwZWVhOGE3MDZjNGMzNGExNjg5MWY4NGU3YiIsImV4cCI6MTU4MzA2MDEzNSwiaXNzIjoidGVzdC1ibG9nIn0.80GcIbMCMzzgWU0KJeU_4m1yhpgZz6vfoUiVphNHp6A`
			c, err := ParseToken(jwtStr)
			So(err, ShouldBeNil)
			So(c.Password, ShouldEqual, EncodeMD5("12345"))
			So(c.Username, ShouldEqual, EncodeMD5("admin"))
		})
	})

}
