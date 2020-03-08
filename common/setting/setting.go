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

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var (
	DefaultServerSetting = &Server{}
	DefaultDBSetting     = &Database{}
	DefaultRedisSetting  = &Redis{}

	cfg *ini.File
	err error
)

func GetDatabaseConf() *Database {
	return DefaultDBSetting
}

func GetServerConf() *Server {
	return DefaultServerSetting
}

func GetRedisConf() *Redis {
	return DefaultRedisSetting
}

func Setup(path string) {

	// path = `conf/app.ini`
	cfg, err = ini.Load(path)
	if err != nil {
		log.Fatalf("setting fail: %v", err)
	}

	mapTo("server", DefaultServerSetting)
	mapTo("database", DefaultDBSetting)
	mapTo("redis", DefaultRedisSetting)

	log.Info("== setup conf ==")
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("setting map fail: %v", err)
	}
}
