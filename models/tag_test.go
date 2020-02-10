package models

import "testing"

func TestAddTag(t *testing.T) {
	setUp()

	if err := InitModel(); err != nil {
		t.Error(err)
	}

	tests := []struct {
		Name      string
		State     int
		CreatedBy string
	}{
		{
			Name:      "哲学♂",
			State:     1,
			CreatedBy: "jim",
		},
		{
			Name:      "放松",
			State:     1,
			CreatedBy: "lee",
		},
		{
			Name:      "宅",
			State:     1,
			CreatedBy: "jim",
		},
	}

	for _, v := range tests {
		if err := AddTag(v.Name, v.State, v.CreatedBy); err != nil {
			t.Error(err)
		}
	}

}
