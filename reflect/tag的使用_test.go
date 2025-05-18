package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type person struct {
	Name string `tag:"name"`
	Id   int    `tag:"id"`
}

func TestName(t *testing.T) {
	p := person{Name: "www", Id: 1}

	v := reflect.ValueOf(p)
	tp := reflect.TypeOf(p)
	fmt.Println(v)
	fmt.Println(tp)

	for i := 0; i < tp.NumField(); i++ {
		field := tp.Field(i)          // Field 返回字段信息
		val := v.Field(i).Interface() // 获取字段值
		fmt.Printf("Field: %s, Value: %v, Tag: %s\n", field.Name, val, field.Tag.Get("tag"))
	}

}
