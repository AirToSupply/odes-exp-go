//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"reflect"
)

func reflectValue(value interface{}) {
	fmt.Printf("type: %v, value: %v\n", reflect.TypeOf(value), reflect.ValueOf(value))
}

func main() {
	reflectValue(18)
	reflectValue("string")
	reflectValue(3.141592653)
	reflectValue(true)
}
