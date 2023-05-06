//go:build ignore
// +build ignore

package main

import "fmt"

// define struct
type Book struct {
	title string
	auth  string
}

// [TIPS] transfer by value
func updateValue(book Book) {
	book.auth = "US"
}

// [TIPS] transfer by reference
func updateReference(book *Book) {
	book.auth = "US"
}

func main() {
	var book Book
	book.title = "Golang"
	book.auth = "Zh"
	fmt.Printf("[update value before] book value: %v\n", book)
	updateValue(book)
	fmt.Printf("[update value after] book value: %v\n", book)

	fmt.Printf("[update reference before] book value: %v\n", book)
	updateReference(&book)
	fmt.Printf("[update reference after] book value: %v\n", book)

}
