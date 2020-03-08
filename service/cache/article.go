package cache

import (
	"fmt"
	"sisyphus/common/ecode"
	"strconv"
	"strings"
)

type Article struct {
	ID    int
	TagID int
	State int

	PageNum  int
	PageSize int
}

func (a *Article) GetArticleKey() string {
	return fmt.Sprintf("%s_%d", ecode.CACHE_ARTICLE, a.ID)
}

func (a *Article) GetArticlesKey() string {
	keys := []string{
		ecode.CACHE_ARTICLE,
		"LIST",
	}

	if a.ID > 0 {
		keys = append(keys, strconv.Itoa(a.ID))
	}

	if a.TagID > 0 {
		keys = append(keys, strconv.Itoa(a.TagID))
	}
	if a.State > 0 {
		keys = append(keys, strconv.Itoa(a.State))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}
	return strings.Join(keys, "_")

}
