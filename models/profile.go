package models

import "github.com/jinzhu/gorm"
import . "sisyphus/models/po"

func GetProfileById(id string) (*Profile, error) {
	var profile Profile
	err := db.Where("id = ? AND deleted_on = ?", id, 0).First(&profile).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &profile, nil
}

func EditProfile(id string, data interface{}) error {
	if err := db.Model(&Profile{}).Where("id = ? AND deleted_on = ? ", id, 0).
		Updates(data).Error; err != nil {
		return err
	}

	return nil
}
