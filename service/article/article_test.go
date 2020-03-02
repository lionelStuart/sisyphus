package article

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

func TestArticle_Add(t *testing.T) {

	Convey("setup", t, func() {
		setupSuite()

		Convey("test add article", func() {
			a := Article{
				TagID:         4,
				Title:         "test article add",
				Desc:          "nil",
				Content:       "this is a test",
				CreatedBy:     "admin",
				CoverImageUrl: "blank",
				State:         0,
			}
			err := a.Add()
			So(err, ShouldBeNil)
		})

	})
}

func TestArticle_GetAll(t *testing.T) {
	Convey("setup", t, func() {
		setupSuite()

		Convey("test add article", func() {
			articles := []Article{
				{
					PageNum:  0,
					PageSize: 2,
					TagID:    2,
				},
				{
					PageNum:  0,
					PageSize: 2,
					TagID:    -1,
				},
			}

			for _, val := range articles {
				article, err := val.GetAll()
				So(err, ShouldBeNil)
				t.Logf("recv article: %#+v \n", article)
			}

		})
	})
}
