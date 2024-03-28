package redis

import "encoding/json"

func TojsonStr(i interface{}) string {
	marshal, _ := json.Marshal(i)
	return string(marshal)
}
