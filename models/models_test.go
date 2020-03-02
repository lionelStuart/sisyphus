package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"sisyphus/common/setting"
	"testing"
)

func setupSuite() {
	path := `D:\PHOENIX\Documents\WORKSPACE\Go\GO_WORKSPACE\sisyphus\conf\app.ini`
	setting.Setup(path)
	Setup()
}

func TestExistArticleByID(t *testing.T) {

	Convey("setup", t, func() {
		setupSuite()

		Convey("test not exist article", func() {
			ok, err := ExistArticleByID(101)
			So(ok, ShouldBeFalse)
			So(err, ShouldBeNil)
		})

	})

}

func TestAddArticle(t *testing.T) {
	Convey("setup", t, func() {
		setupSuite()

		Convey("test add article", func() {
			a := map[string]interface{}{
				"tag_id":          2,
				"title":           "test article title",
				"desc":            "desc",
				"content":         "content",
				"created_by":      "jim",
				"cover_image_url": "blank",
				"state":           0,
			}

			err := AddArticle(a)
			So(err, ShouldBeNil)
		})

	})
}

func TestAddTag(t *testing.T) {
	Convey("setup", t, func() {
		setupSuite()

		Convey("test add tag", func() {
			in := []struct {
				name     string
				state    int
				createBy string
			}{
				{
					name:     "sport",
					state:    0,
					createBy: "admin",
				},
				{
					name:     "movie",
					state:    0,
					createBy: "admin",
				},
			}

			for _, i := range in {
				err := AddTag(i.name, i.state, i.createBy)
				So(err, ShouldBeNil)
			}
		})

	})
}
