package models

import "testing"

func TestAddArticle(t *testing.T) {
	setUp()

	tests := map[string]interface{}{
		"tag_id":          1,
		"title":           "test title",
		"desc":            "this is a test title",
		"content":         "this is a test content",
		"created_by":      "jim",
		"state":           0,
		"cover_image_url": "null",
	}

	if err = AddArticle(tests); err != nil {
		panic(err)
	}

}

func TestGetArticle(t *testing.T) {
	setUp()

	id := 1
	article, err := GetArticle(id)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v \n", article)
}

func TestDeleteArticle(t *testing.T) {
	setUp()

	id := 1
	if err := DeleteArticle(id); err != nil {
		t.Error(err)
	}

}

func TestCleanAllArticle(t *testing.T) {
	setUp()

	if err := CleanAllArticle(); err != nil {
		t.Error(err)
	}

}
