package pcpsd

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

//StartHTTPServer 启动http服务
func StartHTTPServer(fd *os.File) {
	e := echo.New()
	e.Use(middleware.Recover())
	//注册日志中间件 默认os.Stdout 也就是会在控制台打印日志
	// e.Use(middleware.Logger())

	i := getWriter("runtime/logs/pcps-2020-02-17.log")
	//修改中间件配置 将日志内容输出到文件
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		//格式化日志 暂时记录全部内容
		// Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: i,
	}))
	//注册路由
	e.GET("/hello", hello)
	e.GET("/ws", startWebsocketServer)
	e.GET("/list", list)
	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, bitch!")
}

func list(c echo.Context) error {
	return c.String(http.StatusOK, "dev_list")
}

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		filename+".%Y-%m-%d"+".log", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}
	return hook
}
