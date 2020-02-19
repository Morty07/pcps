package main

import (
	"os"
	"pcps/internal/logging"
	"pcps/internal/setting"
	"pcps/pcpsd"
)

var file *os.File

func init() {
	//初始化配置
	setting.Setup()
	logging.SetLog()
}

func main() {
	writerLog := logging.GetWriter()
	//启动http服务
	pcpsd.StartHTTPServer(writerLog)
}
