package setting

import (
	"log"
	"pcps/pcpsd/common"
	"time"

	"github.com/go-ini/ini"
)

type PCPS struct {
	PageSize int
	// PrefixUrl string
	JwtSecret string

	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

	CuttingLogInterval time.Duration
	DeleteLogInterval  time.Duration
}

type Server struct {
	RunMode  string
	HttpPort string
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
	PCPSSetting     = &PCPS{}     //项目配置
	ServerSetting   = &Server{}   //服务相关配置
	DatabaseSetting = &Database{} //数据库配置
	RedisSetting    = &Redis{}    //redis配置
	cfg             *ini.File
)

//Setup 初始化配置
func Setup() {
	var err error
	if cfg, err = ini.Load(common.APP_INI_PATH); err != nil {
		log.Fatalf("setting.Setup, fail to parse "+common.APP_INI_PATH+": %v", err)
	}

	mapTo("app", PCPSSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	// mapTo("redis", RedisSetting)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	if err := cfg.Section(section).MapTo(v); err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
