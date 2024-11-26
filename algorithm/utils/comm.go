package utils

import "encoding/json"

func Tojson(v interface{}) string {
    str, _ := json.Marshal(v)
    return string(str)
}

func Min(i, j int) int {
    if i < j {
        return i
    }
    return j
}

func Max(i, j int) int {
    if i > j {
        return i
    }
    return j
}

