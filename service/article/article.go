package article

import (
	"encoding/json"
	"github.com/prometheus/common/log"
	"sisyphus/common/redis"
	"sisyphus/models"
)
import cacheService "sisyphus/service/cache"

type Article struct {
	ID            int
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string

	PageNum  int
	PageSize int
}

func (a *Article) Add() error {
	article := map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"created_by":      a.CreatedBy,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
	}

	if err := models.AddArticle(article); err != nil {
		return err
	}
	return nil
}

func (a *Article) Edit() error {
	return models.EditArticle(a.ID, map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
		"modified_by":     a.ModifiedBy,
	})
}

func (a *Article) Get() (*models.Article, error) {
	// use cache
	cacheSvc := cacheService.Article{ID: a.ID}
	key := cacheSvc.GetArticleKey()
	if exist, _ := redis.DefaultConn.Exists(key); exist {
		data, err := redis.DefaultConn.Get(key)
		if err != nil {
			log.Info(err)
		} else {
			var cache *models.Article
			json.Unmarshal(data, &cache)
			return cache, nil
		}
	}
	// use db
	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}

	//once get article by db, cache it
	err = redis.DefaultConn.Set(key, article, 3000)
	if err != nil {
		//do nothing
	}

	return article, nil
}

func (a *Article) GetAll() ([]models.Article, error) {
	// use cache
	cacheSvc := cacheService.Article{
		TagID:    a.TagID,
		State:    a.State,
		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}
	key := cacheSvc.GetArticlesKey()
	if exist, _ := redis.DefaultConn.Exists(key); exist {
		data, err := redis.DefaultConn.Get(key)
		if err != nil {
			log.Info(err)
		} else {
			var cache []models.Article
			json.Unmarshal(data, &cache)
			return cache, nil
		}
	}
	// use db
	articles, err := models.GetArticles(a.PageNum, a.PageSize, a.GetMaps())
	if err != nil {
		return nil, err
	}
	//once get article by db, cache it
	err = redis.DefaultConn.Set(key, articles, 3000)
	if err != nil {
		//do nothing
	}

	return articles, nil
}

func (a *Article) Delete() error {
	return models.DeleteArticle(a.ID)
}

func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticleByID(a.ID)
}

func (a *Article) Count() (int, error) {
	return models.GetArticleTotal(a.GetMaps())
}

func (a *Article) GetMaps() map[string]interface{} {
	maps := map[string]interface{}{
		"deleted_on": 0,
	}
	if a.State != -1 {
		maps["state"] = a.State
	}
	if a.TagID != -1 {
		maps["tag_id"] = a.TagID
	}

	return maps
}
