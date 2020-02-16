package conver

import (
	"fmt"
)

//InterfaceToRealType 转成本来的类型
func InterfaceToRealType(m map[string]interface{}) interface{} {
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			return v.(string)
		case float64:
			fmt.Println(k, "is float", int64(vv))
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		case nil:
			fmt.Println(k, "is nil", "null")
		case map[string]interface{}:
			fmt.Println(k, "is an map:")
			Print_json(vv)
		default:
			fmt.Println(k, "is of a type I don't know how to handle ", fmt.Sprintf("%T", v))
		}
	}
}
