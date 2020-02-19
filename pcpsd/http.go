package pcpsd

import (
	"io"
	"net/http"
	"pcps/internal/conf"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//StartHTTPServer 启动http服务
func StartHTTPServer(w io.Writer) {
	e := echo.New()
	e.Use(middleware.Recover())
	//注册日志中间件 默认os.Stdout 也就是会在控制台打印日志
	// e.Use(middleware.Logger())
	//修改中间件配置 将日志内容输出到文件
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		//格式化日志 暂时记录全部内容
		// Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: w,
	}))
	//注册路由
	e.GET("/hello", hello)
	e.GET("/ws", startWebsocketServer)
	e.GET("/list", list)
	e.Logger.Fatal(e.Start(":" + conf.GetString("port")))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, bitch!")
}

func list(c echo.Context) error {
	return c.String(http.StatusOK, "dev_list")
}
