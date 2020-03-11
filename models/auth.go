package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"sisyphus/common/utils"
	"strings"
)

type Auth struct {
	Model

	//ID       int    `gorm:"primary_key" json:"id"`
	Uid      string `json:"uid"` //base32
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	State    int    `json: "state"`

	// ProfileID  int `json:"profile_id" gorm:"index"`
	Profile Profile `json:"profile"`
}

func CheckAuth(username, password string) (bool, error) {
	var auth Auth

	err := db.Select("id").Where(Auth{Username: username, Password: password}).
		First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if auth.ID > 0 {
		return true, nil
	}
	return false, nil
}

func AddAuth(data map[string]interface{}) error {
	var auth Auth
	if err := utils.MapToStruct(data, &auth); err != nil {
		return err
	}
	auth.Uid = utils.GenBase32()

	if err := db.Create(&auth).Error; err != nil {
		return err
	}

	auth.Profile.ID = auth.ID
	if err := db.Create(&auth.Profile).Error; err != nil {
		return err
	}

	return nil
}

func GetAuthById(id int) (*Auth, error) {
	var auth Auth
	err := db.Where("id = ? AND deleted_on = ?", id, 0).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &auth, nil
}

func GetProfileById(id int) (*Profile, error) {
	var profile Profile
	err := db.Where("id = ? AND deleted_on = ?", id, 0).First(&profile).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &profile, nil
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
	if auth.ID > 0 {
		return true, nil
	}
	return false, nil
}
