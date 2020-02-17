package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type PCPS struct {
	PageSize  int
	PrefixUrl string
	JwtSecret string

	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var PCPSSetting = &PCPS{}

type Server struct {
	RunMode  string
	HttpPort int
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("pcpsd/conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'pcpsd/conf/app.ini': %v", err)
	}

	mapTo("app", PCPSSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
