//go:build ignore
// +build ignore

package main

import "fmt"

func foo(object interface{}) {
	fmt.Printf("invoke value type : %T, value : %v\n", object, object)
}

type Book struct {
	Name string
}

// interface data type is equivalent to `Object` in Java
func main() {
	foo(Book{Name: "Golang"})
	foo(100)
	foo("abc")
	foo(3.1415926)
}
