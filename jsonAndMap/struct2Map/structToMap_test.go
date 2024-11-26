package struct2Map

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"testing"
	"unsafe"
)

func TestName(t *testing.T) {
	ip := ipcInfo{
		id:       1,
		uuid:     "222",
		ip:       "2",
		deviceId: "22",
	}
	toMap, err := structToMap(&ip)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(len(toMap))
	fmt.Println(toMap)

	marshal, err := json.Marshal(&a{"xxx"})
	fmt.Println(string(marshal))

}

type a struct {
	Name string
}

func structToMap(s interface{}) (map[string]interface{}, error) {
	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	// 传入的是指针
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct but got %s", typ.Kind())
	}

	result := make(map[string]interface{})
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name

		// 只能处理导出的字段
		if field.CanInterface() {
			result[fieldName] = field.Interface()
		} else {
			// 处理未导出的字段，用unsafe处理未导出的字段
			fieldPtr := unsafe.Pointer(field.UnsafeAddr())
			fieldVal := reflect.NewAt(field.Type(), fieldPtr).Elem().Interface()
			result[fieldName] = fieldVal
		}

	}
	return result, nil
}

type ipcInfo struct {
	id               int64
	uuid             string
	ip               string
	deviceId         string
	appearance       string
	name             string
	tp               string
	proto            string
	ptzType          string
	vendor           string
	username         string
	password         string
	controlPoint     string
	model            string
	defaultCoding    string
	nvrIp            string
	elevation        string
	latitude         string
	longitude        string
	workingStatus    string
	onlineStatus     string
	relayStatus      string
	storageStatus    string
	alarmStatus      string
	ipcStatus        string
	platformDeviceId string
	civilId          int64
	civilUuid        string
	organizationId   int64
	gbUuid           string // 共享时的uuid
}

var civilIdRegExp = regexp.MustCompile(`^\d{2,8}$`)

func TestNamel(t *testing.T) {
	i, _ := strconv.ParseFloat("1.23546213", 64)
	fmt.Println(i)
}
