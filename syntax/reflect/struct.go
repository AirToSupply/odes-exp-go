//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int    `info:"id" doc:"标识"`
	Name string `info:"name" doc:"姓名"`
	Age  int    `info:"age" doc:"性别 "`
}

func (this User) Call() {
	fmt.Printf("Call invoke value %v\n", this)
}

func parseStruct(value interface{}) {
	// Parse Type
	valueType := reflect.TypeOf(value)
	fmt.Println("[Parse Type]")
	fmt.Printf("Type: %v\n", valueType)

	// Parse Value
	v := reflect.ValueOf(value)
	fmt.Println("[Parse Value]")
	fmt.Printf("Value: %v\n", v)

	// Parse Field
	fmt.Println("[Parse Field]")
	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		value := v.Field(i).Interface()
		fmt.Printf("%s : %v = %v\n", field.Name, field.Type, value)
	}

	// Parse Method
	fmt.Println("[Parse Method]")
	for i := 0; i < valueType.NumMethod(); i++ {
		method := valueType.Method(i)
		fmt.Printf("%s : %v\n", method.Name, method.Type)
	}
}

func parseStructTag(value interface{}) {
	fmt.Println("[Parse Tag]")
	t := reflect.TypeOf(value).Elem()
	for i := 0; i < t.NumField(); i++ {
		info := t.Field(i).Tag.Get("info")
		doc := t.Field(i).Tag.Get("doc")
		fmt.Printf("info: %v, doc: %v\n", info, doc)
	}
}

func main() {
	user := User{Id: 1, Name: "AirToSUpply", Age: 18}
	parseStruct(user)
	parseStructTag(&user)
}
