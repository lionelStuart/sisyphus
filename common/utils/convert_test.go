package utils

import (
	"sisyphus/common/app"
	"testing"
)
import . "github.com/smartystreets/goconvey/convey"

type Sample struct {
	Name string
	ID   int64
}

type SampleB struct {
	Age int
	Sample
}

type SampleC struct {
	NickName string `mapstructure:"name"`
	UID      int    `mapstructure:"id"`
}

func TestMapToStruct(t *testing.T) {

	Convey("setup", t, func() {
		Convey("test simple", func() {
			from := app.H{
				"name": "jimmy",
				"id":   13,
			}
			var s Sample
			err := MapToStruct(from, &s)
			So(err, ShouldBeNil)
			So(s.ID, ShouldEqual, from["id"])
			So(s.Name, ShouldEqual, from["name"])
		})

		Convey("test lack", func() {
			from := app.H{
				"name": "jimmy",
			}
			var s Sample
			err := MapToStruct(from, &s)
			So(err, ShouldBeNil)
			So(s.Name, ShouldEqual, from["name"])
		})

		Convey("test nesting", func() {
			from := app.H{
				"sample": app.H{
					"name": "jimmy",
					"id":   13,
				},
				"age": 15,
			}
			var s SampleB
			err := MapToStruct(from, &s)
			So(err, ShouldBeNil)
			So(s.ID, ShouldEqual, from["sample"].(map[string]interface{})["id"])
			So(s.Name, ShouldEqual, from["sample"].(map[string]interface{})["name"])
			So(s.Age, ShouldEqual, from["age"])
		})

		Convey("test annotation", func() {
			from := app.H{
				"name": "jimmy",
				"id":   13,
			}
			var s SampleC
			err := MapToStruct(from, &s)
			So(err, ShouldBeNil)
			So(s.UID, ShouldEqual, from["id"])
			So(s.NickName, ShouldEqual, from["name"])
		})
	})

}
