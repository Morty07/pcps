package pcpsd

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

//ParseLogCameraInfo 解析日志中摄像头传递的信息
func ParseLogCameraInfo(s string) (string, map[string]interface{}, error) {
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

//FilterByConditionCameraInfo 筛选解析出的信息
func FilterByConditionCameraInfo() {

}
