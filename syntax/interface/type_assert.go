//go:build ignore
// +build ignore

package main

import "fmt"

func typeAssert(object interface{}) {
	value, ok := object.(string)
	if !ok {
		fmt.Printf("object is not string type\n")
	} else {
		fmt.Printf("object is string type\n")
		fmt.Printf("invoke value type : %T, value : %v\n", value, value)
	}
}

func OfSomeType(object interface{}) {
	switch t := object.(type) {
	default:
		fmt.Printf("Unexpected type %T, value: %v\n", t, t)
	case bool:
		fmt.Printf("Boolean type %T, value: %v\n", t, t)
	case string:
		fmt.Printf("String type %T, value: %v\n", t, t)
	case int:
		fmt.Printf("Int type %T, value: %v\n", t, t)
	case float64:
		fmt.Printf("Float type %T, value: %v\n", t, t)
	case *Book:
		fmt.Printf("Pointer Of Book type %T, value: %v\n", t, t)
	}
}

type Book struct {
	Name string
}

// By using interface types, the underlying data types can be distinguished.
func main() {
	// type assert
	typeAssert(Book{Name: "Golang"})
	typeAssert(100)
	typeAssert("abc")
	typeAssert(3.1415926)

	// of some type
	OfSomeType(Book{Name: "Golang"})
	OfSomeType(100)
	OfSomeType("abc")
	OfSomeType(3.1415926)
	OfSomeType(&Book{Name: "Golang"})
}
