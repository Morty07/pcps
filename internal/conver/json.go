package conver

import (
	"encoding/json"
	"fmt"
)

func JsonToMap(j string) (map[string]interface{}, error) {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(j), &m)
	if err != nil {
		return nil,
			fmt.Errorf("json to map err: %v", err)
	}
	return m, nil
}
