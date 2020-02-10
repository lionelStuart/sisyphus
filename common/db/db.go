package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Conf struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}

func NewDB(dbConf Conf) (*gorm.DB, error) {

	db, err := gorm.Open(dbConf.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConf.User,
		dbConf.Password,
		dbConf.Host,
		dbConf.Name))
	return db, err
}
