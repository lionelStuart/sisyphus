package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Uid  int    `gorm:"unique_index"`
	Name string `json:"name"`
}

func AddCategory(category *Category) error {
	//if err := DB.Create(category).Error; err != nil {
	//	return err
	//}

	return nil
}
