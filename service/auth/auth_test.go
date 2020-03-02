package auth

import (
	. "github.com/smartystreets/goconvey/convey"
	"sisyphus/common/setting"
	"sisyphus/models"
	"testing"
)

func setupSuite() {
	path := `D:\PHOENIX\Documents\WORKSPACE\Go\GO_WORKSPACE\sisyphus\conf\app.ini`
	setting.Setup(path)
	models.Setup()
}

func TestAuth_Check(t *testing.T) {

	Convey("setup", t, func() {
		setupSuite()

		Convey("test auth", func() {
			auth := Auth{
				Username: "master",
				Password: "test123",
			}
			ok, err := auth.Check()
			So(err, ShouldBeNil)
			So(ok, ShouldBeTrue)

		})

	})
}
