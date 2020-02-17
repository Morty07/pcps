package pcpsd

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//StartHTTPServer 启动http服务
func StartHTTPServer(fd *os.File) {
	e := echo.New()
	e.Use(middleware.Recover())
	//注册日志中间件 默认os.Stdout 也就是会在控制台打印日志
	e.Use(middleware.Logger())
	//修改中间件配置 将日志内容输出到文件
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		// Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: fd,
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
