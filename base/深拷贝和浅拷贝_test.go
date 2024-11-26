package base

import (
	"fmt"
	"testing"
)

// 关于对象的拷贝，有这些工具库可以用 github.com/jinzhu/copier

/**


浅拷贝 (Shallow Copy)

如果字段是基本类型（如整数、浮点数、布尔值等），这些值会被直接复制。
如果字段是引用类型（如数组、切片、指针、结构体等），浅拷贝会复制这些引用，而不是引用的实际对象。

浅拷贝，原对象和新对象共享引用的对象。

*/

func TestShallow(t *testing.T) {
	addr := &Address{City: "San Francisco", State: "CA"}
	person1 := &Person{Name: "Alice", Age: 30, Address: addr}

	// 浅拷贝
	person2 := *person1

	// 修改 person2 的 Address
	person2.Address.City = "Los Angeles"

	// person1 的 Address 也会受影响
	fmt.Println(person1.Address.City) // 输出: Los Angeles
	fmt.Println(person2.Address.City) // 输出: Los Angeles
}

/*
*
深拷贝 (Deep Copy)

新对象和原对象完全独立。

所有字段的值都会被复制，包括引用类型指向的对象也会被递归地复制。
深拷贝后，原对象和新对象不会共享引用的对象，它们是完全独立的。
*/
func TestDeep(t *testing.T) {
	addr := &Address{City: "San Francisco", State: "CA"}
	person1 := &Person{Name: "Alice", Age: 30, Address: addr}

	// 深拷贝
	person2 := DeepCopyPerson(person1)

	// 修改 person2 的 Address
	person2.Address.City = "Los Angeles"

	// person1 的 Address 不会受影响
	fmt.Println(person1.Address.City) // 输出: San Francisco
	fmt.Println(person2.Address.City) // 输出: Los Angeles
}

type Person struct {
	Name    string
	Age     int
	Address *Address
}

type Address struct {
	City  string
	State string
}

func DeepCopyAddress(addr *Address) *Address {
	if addr == nil {
		return nil
	}
	return &Address{
		City:  addr.City,
		State: addr.State,
	}
}

func DeepCopyPerson(p *Person) *Person {
	if p == nil {
		return nil
	}
	return &Person{
		Name:    p.Name,
		Age:     p.Age,
		Address: DeepCopyAddress(p.Address),
	}
}

/*
*
使用工具库来深浅拷贝, 比如 github.com/jinzhu/copier
*/
func TestCopier(t *testing.T) {

}
