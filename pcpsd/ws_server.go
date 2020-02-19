package pcpsd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pcps/pcpsd/common"

	// "pcps/internal/file"
	"time"

	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

var rateLimiter = time.Tick(time.Second * 30)

type Line struct {
	ImportStr string
	Parser    string
}

type ParseResult struct {
	Lines    []Line
	LineNum  int
	LastTime int64
}

func startWebsocketServer(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			<-rateLimiter
			result := ParseResult{}
			// file.ReadFileLineByLine()
			m, err := ReadFileByLine(result)
			if err != nil {
				log.Fatal(err)
			}
			err = websocket.Message.Send(ws, m)
			if err != nil {
				log.Fatal(err)
			}

			// msg := ""
			// err = websocket.Message.Receive(ws, &msg)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// fmt.Printf("%s\n", msg)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func ReadFileByLine(r ParseResult) (string, error) {
	f, err := os.Open(common.READ_LOG_PATH)
	if err != nil {
		return "", err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	line := 1

	for scanner.Scan() {
		str := scanner.Text()

		s, err := ParseLogCameraInfo(r, str)
		if err != nil {
			return "", err
		}
		fmt.Println(s)
		time.Sleep(time.Second * 1)
		line++
	}

	if err := scanner.Err(); err != nil {
		// Handle the error
		return "", err
	}
	return "", nil
}
