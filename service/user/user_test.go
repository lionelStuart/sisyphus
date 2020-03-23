package user

import (
	. "github.com/smartystreets/goconvey/convey"
	"sisyphus/common/redis"
	"sisyphus/common/setting"
	"sisyphus/models"
	"testing"
)

func setupSuite() {
	path := `D:\PHOENIX\Documents\WORKSPACE\Go\GO_WORKSPACE\sisyphus\conf\app.ini`
	setting.Setup(path)
	models.Setup()
	redis.SetUp()
}

func TestUser_Add(t *testing.T) {
	Convey("setup", t, func() {
		setupSuite()

		Convey("test add user", func() {
			user := User{
				Username: "jim",
				Password: "pass123",
				Email:    "jim@126.com",
				Phone:    "12300001111",
				State:    0,
				Profile: Profile{
					Nickname: "jim nick",
					Age:      10,
					Gender:   "M",
					Address:  "this is address",
				},
			}

			err := user.Add()
			So(err, ShouldBeNil)
		})
	})
}

func TestUser_EditProfile(t *testing.T) {
	Convey("setup", t, func() {
		setupSuite()

		Convey("test add user", func() {
			user := User{
				ID: "bnxfujkceyryy",
				Profile: Profile{
					Nickname: "jim nick",
					Age:      10,
					Gender:   "F",
					Address:  "this is address",
				},
			}

			err := user.EditProfile()
			So(err, ShouldBeNil)
		})
	})
}
