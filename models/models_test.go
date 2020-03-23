package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"sisyphus/common/app"
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
			a := app.H{
				"tag_id":          3,
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

func TestAddAuth(t *testing.T) {
	Convey("setup", t, func() {
		setupSuite()

		Convey("test add auth", func() {
			a := app.H{
				"username": "tom",
				"password": "pass",
				"email":    "tom@126.com",
				"phone":    "130123412344",
				"state":    0,
				"profile": app.H{
					"nickname": "big tom",
					"age":      15,
					"gender":   "M",
					"address":  "king street",
				},
			}

			err := AddAuthProfile(a)
			So(err, ShouldBeNil)
		})

	})
}

func TestGetAuthOrProfile(t *testing.T) {
	Convey("setup", t, func() {
		setupSuite()

		Convey("test get auth", func() {
			id := "1"
			auth, err := GetAuthById(id)
			So(err, ShouldBeNil)
			So(auth.ID, ShouldEqual, id)
			So(auth.Profile.ID, ShouldEqual, "")
			t.Logf("%+v", auth)
		})

		Convey("test get profile", func() {
			id := "1"
			profile, err := GetProfileById(id)
			So(err, ShouldBeNil)
			So(profile.ID, ShouldEqual, id)
			t.Logf("%+v", profile)
		})
	})
}

func TestDeleteAuthProfile(t *testing.T) {
	Convey("setup", t, func() {
		setupSuite()

		Convey("test get auth", func() {
			id := 4
			err := DeleteAuthProfile(id)
			So(err, ShouldBeNil)
		})
	})
}

func TestExistAuthCondition(t *testing.T) {
	Convey("setup", t, func() {
		setupSuite()

		Convey("test one", func() {
			cond := app.H{
				"id": 7,
			}
			exist, err := ExistAuthCondition(cond)
			So(err, ShouldBeNil)
			So(exist, ShouldBeTrue)
		})

		Convey("test false", func() {
			cond := app.H{
				"id": -100,
			}
			exist, err := ExistAuthCondition(cond)
			So(err, ShouldBeNil)
			So(exist, ShouldBeFalse)
		})

		Convey("test two", func() {
			cond := app.H{
				"id":       -100,
				"username": "master",
			}
			exist, err := ExistAuthCondition(cond)
			So(err, ShouldBeNil)
			So(exist, ShouldBeTrue)
		})

		Convey("test three", func() {
			cond := app.H{
				"id":       -100,
				"username": "tom",
				"phone":    "tom@126.com",
			}
			exist, err := ExistAuthCondition(cond)
			So(err, ShouldBeNil)
			So(exist, ShouldBeTrue)
		})

	})
}
