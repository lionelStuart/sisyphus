package models

import (
	"testing"
)

func TestInitModel(t *testing.T) {
	if err := InitModel(); err != nil {
		t.Error(err)
	}
}

func setUp() {
	if err := InitModel(); err != nil {
		panic(err)
	}
}

func TestAddCategory(t *testing.T) {
	if err := InitModel(); err != nil {
		t.Error(err)
	}

	tests := []*Category{
		{
			Uid:  1,
			Name: `娱乐`,
		},
		{
			Uid:  2,
			Name: `教育`,
		},
		{
			Uid:  3,
			Name: `科技`,
		},
	}
	for _, v := range tests {
		if err := AddCategory(v); err != nil {
			t.Error(err)
		}
	}

}
