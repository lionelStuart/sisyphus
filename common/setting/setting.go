package setting

import (
	"github.com/go-ini/ini"
	"github.com/prometheus/common/log"
	"time"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var (
	DefaultServerSetting = &Server{}
	DefaultDBSetting     = &Database{}

	cfg *ini.File
	err error
)

func GetDatabaseConf() *Database {
	return DefaultDBSetting
}

func GetServerConf() *Server {
	return DefaultServerSetting
}

func Setup(path string) {

	// path = `conf/app.ini`
	cfg, err = ini.Load(path)
	if err != nil {
		log.Fatalf("setting fail: %v", err)
	}

	mapTo("server", DefaultServerSetting)
	mapTo("database", DefaultDBSetting)

	log.Info("== setup conf ==")
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("setting map fail: %v", err)
	}
}
