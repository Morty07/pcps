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
	file = logging.Setup()
}

func main() {
	//启动http服务
	pcpsd.StartHTTPServer(file)
}
