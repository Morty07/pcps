package pcpsd

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//ParseLogCameraInfo 解析日志中摄像头传递的信息
func ParseLogCameraInfo(r *ParseResult, s string) (bool, error) {
	startIndex := strings.Index(s, "{")
	endIndex := strings.LastIndex(s, "}")
	json := s[startIndex : endIndex+1]
	//这里好像只能json to map 因为json内的字段是动态的 转成struct会丢失字段
	// mapResult, err := conver.JsonToMap(json)
	// if err != nil {
	// 	return false, err
	// }
	cameraRe := `log_date - [0-9/]*, [0-9:0-9:0-9 A-Z]* : ([^ :]+)`
	if isMatches, _ := regexp.MatchString(cameraRe, s); !isMatches {
		return false, fmt.Errorf("wrong match string: %v", s)
	}
	cameraRegexp := regexp.MustCompile(cameraRe)
	matches := cameraRegexp.FindAllStringSubmatch(s, -1)
	return FilterByConditionCameraInfo(r, json, matches[0][1])
}

//FilterByConditionCameraInfo 筛选解析出的信息
func FilterByConditionCameraInfo(r *ParseResult, m string, s string) (bool, error) {
	i := strings.Index(m, "\"time\": \"")
	t, err := strconv.ParseInt(m[i+9:i+22], 10, 64)
	if err != nil {
		return false, err
	}
	//以下注释是根据日志中的数据得出来的测试数据
	tenUnixTime := t / 1000
	// fmt.Println(tenUnixTime)
	// now := time.Now().Unix()
	var now int64
	now = 1576052159
	// left := now - 30
	left := now - 120
	if r.LastTime != 0 {
		left = r.LastTime
	}
	if tenUnixTime >= left && tenUnixTime <= now {
		r.Lines = append(r.Lines, Line{
			ImportStr: s,
			Parser:    m,
		})

		return true, nil
	}
	r.LastTime = t
	return false, fmt.Errorf("it's the last one")
}
