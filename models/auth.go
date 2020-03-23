package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"sisyphus/common/utils"
	. "sisyphus/models/po"
	"strings"
)

func CheckAuth(username, password string) (bool, error) {
	var auth Auth

	err := db.Select("id").Where(Auth{Username: username, Password: password}).
		First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	//if auth.ID > 0 {
	//	return true, nil
	//}
	return false, nil
}

func AddAuthProfile(data map[string]interface{}) error {
	var auth Auth
	if err := utils.MapToStruct(data, &auth); err != nil {
		return err
	}
	auth.ID = utils.GenBase32()

	if err := db.Create(&auth).Error; err != nil {
		return err
	}

	auth.Profile.ID = auth.ID
	if err := db.Create(&auth.Profile).Error; err != nil {
		return err
	}

	return nil
}

func GetAuthById(id string) (*Auth, error) {
	var auth Auth
	err := db.Where("id = ? AND deleted_on = ?", id, 0).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &auth, nil
}

func EditAuth(id string, data interface{}) error {
	if err := db.Model(&Auth{}).Where("id = ? AND deleted_on = ? ", id, 0).
		Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func ExistAuthCondition(condition map[string]interface{}) (bool, error) {
	var gather []string
	for k, v := range condition {
		gather = append(gather, fmt.Sprintf("%s = '%v'", k, v))
	}

	var auth Auth
	err := db.Select("id").Where(strings.Join(gather, " OR ")).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if len(auth.ID) > 0 {
		return true, nil
	}
	//if auth.ID > 0 {
	//	return true, nil
	//}
	return false, nil
}

func DeleteAuthProfile(id int) error {
	var err error

	if err = db.Where("id = ? AND deleted_on = ?", id, 0).Delete(Auth{}).Error; err != nil {
		return err
	}

	if err = db.Where("id = ?", id).Delete(Profile{}).Error; err != nil {
		return err
	}

	return nil
}
