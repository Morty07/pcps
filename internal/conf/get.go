package conf

import (
	"pcps/internal/setting"
	"time"
)

//GetString 获取string类型的全局配置文件 这种做法较low 暂时先这样
func GetString(name string) string {
	switch name {
	case "jwt_secret":
		return setting.PCPSSetting.JwtSecret
	case "port":
		return setting.ServerSetting.HttpPort
	default:
		return ""
	}
}

//GetTime 时间类型的配置
func GetTime(name string) time.Duration {
	switch name {
	case "cutting_log_interval":
		return setting.PCPSSetting.CuttingLogInterval
	case "delete_log_interval":
		return setting.PCPSSetting.DeleteLogInterval
	default:
		return 0
	}
}
