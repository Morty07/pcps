package logging

import (
	"fmt"
	"io"
	"pcps/internal/file"
	"pcps/internal/setting"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.PCPSSetting.RuntimeRootPath, setting.PCPSSetting.LogSavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s.%s",
		// setting.PCPSSetting.LogSaveName,
		time.Now().Format(setting.PCPSSetting.TimeFormat),
		setting.PCPSSetting.LogFileExt,
	)
}

func GetWriter() io.Writer {
	absolutePath, err := file.GetAbsolutePath(getLogFilePath())
	if err != nil {
		panic(err)
	}
	// fileName := getLogFileName()
	// file := absolutePath + fileName
	hook, err := rotatelogs.New(
		absolutePath+"%Y-%m-%d.log", // 没有使用go风格反人类的format格式
		// rotatelogs.WithLinkName(file),             //生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Hour*24*7),     //文件最大保存时间
		rotatelogs.WithRotationTime(time.Hour*24), //日志切割时间间隔
	)

	if err != nil {
		panic(err)
	}
	return hook
}
