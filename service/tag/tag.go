package tag

import (
	"encoding/json"
	"github.com/prometheus/common/log"
	"io"
	"sisyphus/common/redis"
	"sisyphus/models"
	"sisyphus/models/po"
	cacheService "sisyphus/service/cache"
)

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

func (t *Tag) ExistByID() (bool, error) {
	return models.ExistTagByID(t.ID)
}

func (t *Tag) Add() error {
	return models.AddTag(t.Name, t.State, t.CreatedBy)
}

func (t *Tag) Edit() error {
	data := map[string]interface{}{
		"modified_by": t.ModifiedBy,
		"name":        t.Name,
	}
	if t.State > 0 {
		data["state"] = t.State
	}

	return models.EditTag(t.ID, data)
}

func (t *Tag) Delete() error {
	return models.DeleteTag(t.ID)
}

func (t *Tag) Count() (int, error) {
	return models.GetTagTotal(t.getMaps())
}

func (t *Tag) GetAll() ([]po.Tag, error) {
	//
	cacheSvc := cacheService.Tag{
		State: t.State,

		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}
	key := cacheSvc.GetTagsKey()
	if exist, _ := redis.DefaultConn.Exists(key); exist {
		data, err := redis.DefaultConn.Get(key)
		if err != nil {
			log.Info(err)
		} else {
			var cache []po.Tag
			json.Unmarshal(data, &cache)
			return cache, nil
		}
	}

	tags, err := models.GetTags(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}
	err = redis.DefaultConn.Set(key, tags, 3600)
	if err != nil {
		//TODO
	}
	return tags, nil
}

func (t *Tag) Export() (string, error) {
	return "", nil
}

func (t *Tag) Import(r io.Reader) error {
	return nil
}

func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0

	if t.Name != "" {
		maps["name"] = t.Name
	}

	if t.State > 0 {
		maps["state"] = t.State
	}
	return maps
}
