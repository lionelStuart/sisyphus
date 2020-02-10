package db

import (
	"testing"
)

var testConf Conf

func init() {
	testConf = Conf{
		Type:     "mysql",
		User:     "root",
		Password: "neon1234",
		Host:     "127.0.0.1:3309",
		Name:     "sisyphus",
	}
}

func TestNewDB(t *testing.T) {
	db, err := NewDB(testConf)
	defer db.Close()
	if err != nil {
		t.Error("fail ", err)
	}
	if err = db.DB().Ping(); err != nil {
		t.Error("fail ", err)
	}

}
