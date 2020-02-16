package pcpsd

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//StartHTTPServer 启动http服务
func StartHTTPServer() {
	e := echo.New()
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	path := "/Users/chenle/Workspaces/Code/company/pcps/runtime/pcpsd/2019-02-17.log"
	fd, _ := os.OpenFile(
		path,
		os.O_RDWR|os.O_APPEND,
		0666,
	)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: fd,
	}))
	e.GET("/hello", hello)
	e.GET("/ws", startWebsocketServer)
	e.GET("/list", list)
	// putLog(e)

	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, bitch!")
}

func list(c echo.Context) error {
	return c.String(http.StatusOK, "dev_list")
}

func putLog(e *echo.Echo) {

	// e.Logger.SetOutput(fd)
}
