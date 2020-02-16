package main

import (
	"fmt"
	"pcps/internal/file"
)

func main() {
	// s, err := file.ParseCameraInfo(`[2020-01-16T18:11:22.540] [INFO] log_date - 1/16/2020, 6:11:22 PM : 3OQt8qoKAnOfl6aOhLi3q9Qh/Eh90SFCy : {"status": "3", "root_mac": "b4e62dfcccdd", "dev_list": [{"mac": "b4e62dfccdad", "park_num": "A123"}], "time": "1579169515015"}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(s)
	// scanner := bufio.NewScanner(
	// 	strings.NewReader("ABCDEFG\nHIJKELM")
	// )
	// for scanner.Scan(){
	// 	fmt.Println(scanner.Text())
	// }
	// bufio.NewReader()
	// fmt.Println(strings.NewReader("ABCDEFG\nHIJKELM"))
	// e := echo.New()
	// now := time.Now().Format("2006-01-02 15:04:05")
	// fmt.Println(now)
	// nowU := time.Now().Unix()
	// fmt.Println(nowU)
	// l, err := file.ReadFile()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(l)
	// str := `log_date - [0-9/]*, [0-9:0-9:0-9 A-Z]* : ([^:]+)`
	// contents := `[2019-12-30T10:06:30.293] [INFO] log_date - 12/30/2019, 10:06:30 AM : 3OQt8qoKAnOfl6aOhLi3q9Qh/ORlhKi8l : {"status": 3, "dev": "b4e62dfccd31", "code": 1, "time": "1577671621765"}`
	// profileRe := regexp.MustCompile(str)
	// matches := profileRe.FindAllStringSubmatch(contents, -1)
	// // a, _ := regexp.MatchString(str, contents)
	// // a, _ := regexp.MatchString(`log_date - [0-9/]*, [0-9:0-9:0-9 A-Z]* : ([^:]+)`, "log_date - 12/30/2019, 10:06:30 AM : 3OQt8qoKAnOfl6aOhLi3q9Qh/ORlhKi8l")
	// fmt.Println(matches[0][1])

	// context := strings.Fields("zhangsan lisi wanger")
	// fmt.Println(context)

	// e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	// // e.Static("/", "../public")
	// e.GET("/ws", pcpsd.StartWebsocketServer)
	// e.Logger.Fatal(e.Start(":1323"))

	_, m, err := file.ParseCameraInfo(`[2019-12-30T10:06:30.293] [INFO] log_date - 12/30/2019, 10:06:30 AM : 3OQt8qoKAnOfl6aOhLi3q9Qh/ORlhKi8l : {"status": 3, "dev": "b4e62dfccd31", "code": 1, "time": "1577671621765"}`)
	if err != nil {
		fmt.Println(err)
	}
	file.Print_json(m)
}
