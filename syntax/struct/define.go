//go:build ignore
// +build ignore

package main

import "fmt"

// define struct
type Book struct {
	title string
	auth  string
}

func main() {
	// Method 1
	var book Book
	book.title = "Golang"
	book.auth = "Zh"
	fmt.Printf("book type: %T, book value: %v\n", book, book)

	// Method 2
	book2 := Book{
		title: "Java",
		auth:  "US",
	}
	fmt.Printf("book2 type: %T, book2 value: %v\n", book2, book2)

}
