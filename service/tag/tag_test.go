package tag

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

func TestExist(t *testing.T) {

	Convey("setup", t, func() {
		setupSuite()

		Convey("test exist by name", func() {
			tag := Tag{
				Name: "sport",
			}
			ok, err := tag.ExistByName()
			So(ok, ShouldBeTrue)
			So(err, ShouldBeNil)
		})

		Convey("test exist by id", func() {
			tag := Tag{
				ID: 2,
			}
			ok, err := tag.ExistByID()
			So(ok, ShouldBeTrue)
			So(err, ShouldBeNil)
		})

	})
}

func TestTag_Add(t *testing.T) {
	Convey("setup", t, func() {
		setupSuite()

		Convey("test add tag", func() {
			tag := Tag{
				Name:      "music",
				CreatedBy: "admin",
				State:     0,
			}
			err := tag.Add()
			So(err, ShouldBeNil)
		})

	})
}

func TestTag_Edit(t *testing.T) {
	Convey("setup", t, func() {
		setupSuite()

		Convey("test edit tag", func() {
			tag := Tag{
				Name:       "music-Fork",
				ModifiedBy: "tom",
				State:      0,
				ID:         4,
			}
			err := tag.Edit()
			So(err, ShouldBeNil)
		})

	})
}

func TestTag_GetAll(t *testing.T) {
	Convey("setup", t, func() {
		setupSuite()

		Convey("test get all tag", func() {
			tags := []Tag{
				{
					PageNum:  0,
					PageSize: 1,
				},
				{
					PageNum:  1,
					PageSize: 2,
				}, {
					PageNum:  1,
					PageSize: 3,
				},
			}

			for _, val := range tags {
				tags, err := val.GetAll()
				So(err, ShouldBeNil)
				t.Logf("recv tag: %#+v \n", tags)
			}

		})

	})
}
