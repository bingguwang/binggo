package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type people struct {
}

func (*people) Say() {
	fmt.Println("hi ")
}
func TestNamexx(t *testing.T) {
	p := people{}
	val := reflect.ValueOf(&p)
	mname := val.MethodByName("Say")
	mname.Call([]reflect.Value{})
}
