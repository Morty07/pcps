package pcpsd

import (
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

var rateLimiter = time.Tick(time.Second * 30)

func StartWebsocketServer(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			<-rateLimiter
			err := websocket.Message.Send(ws, "hello client!")
			if err != nil {
				log.Fatal(err)
			}

			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", msg)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
