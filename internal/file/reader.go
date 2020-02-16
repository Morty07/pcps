package file

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

var filePath = "/Users/chenle/Workspaces/Code/company/pcps/internal/file/temp/camera.log"

// func ReadFile() (int, error) {
// 	f, err := os.Open(filePath)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer f.Close()
// 	scanner := bufio.NewScanner(f)
// 	line := 1

// 	for scanner.Scan() {
// 		str := scanner.Text()
// 		now := 1579169482
// 		// 1579169515015
// 		s, m, err := parseCameraInfo(str)
// 		if err != nil {
// 			return 0, err
// 		}
// 		logTime := 1579169515555 / 1000
// 		nNow := now - 30
// 		if logTime >= nNow && logTime <= now {
// 			// t := strings.Fields(scanner.Text())
// 			// t := strings.Fields("[2020-01-17T02:27:30.452] [INFO] log_date - mysqlerror:Error: Connection lost: The server closed the connection.")

// 			// for k, v := range t {
// 			// 	fmt.Println("ketis", k, "val is", v)
// 			// }
// 			return line, nil
// 		}
// 		fmt.Println(scanner.Text())
// 		time.Sleep(time.Second * 5)
// 		line++
// 	}
// 	return 0, nil
// 	if err := scanner.Err(); err != nil {
// 		// Handle the error
// 	}
// }

//parseCameraInfo 解析摄像头日志中的信息
func ParseCameraInfo(s string) (string, map[string]interface{}, error) {
	startIndex := strings.Index(s, "{")
	endIndex := strings.LastIndex(s, "}")
	infoJSON := s[startIndex : endIndex+1]
	var infoMap map[string]interface{}
	err := json.Unmarshal([]byte(infoJSON), &infoMap)
	if err != nil {
		return "",
			nil,
			fmt.Errorf("json to map err: %v", err)
	}
	cameraRe := `log_date - [0-9/]*, [0-9:0-9:0-9 A-Z]* : ([^ :]+)`
	if isMatches, _ := regexp.MatchString(cameraRe, s); !isMatches {
		return "",
			nil,
			fmt.Errorf("wrong match string: %v", s)
	}
	cameraRegexp := regexp.MustCompile(cameraRe)
	matches := cameraRegexp.FindAllStringSubmatch(s, -1)
	return matches[0][1], infoMap, nil
}

// int1, _ := strconv.Atoi(nv)

// f,err := os.Open(path)
// if err != nil {
//     return 0,err
// }
// defer f.Close()

// // Splits on newlines by default.
// scanner := bufio.NewScanner(f)

// line := 1
// // https://golang.org/pkg/bufio/#Scanner.Scan
// for scanner.Scan() {
//     if strings.Contains(scanner.Text(),"yourstring") {
//         return line,nil
//     }

//     line++
// }

// if err := scanner.Err(); err != nil {
//     // Handle the error
// }
