package logging

import (
	"fmt"
	"pcps/internal/setting"
	"time"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.PCPSSetting.RuntimeRootPath, setting.PCPSSetting.LogSavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.PCPSSetting.LogSaveName,
		time.Now().Format(setting.PCPSSetting.TimeFormat),
		setting.PCPSSetting.LogFileExt,
	)
}
