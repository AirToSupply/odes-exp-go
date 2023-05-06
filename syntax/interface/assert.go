//go:build ignore
// +build ignore

package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read Book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write Book")
}

func main() {
	book := &Book{}

	var r Reader
	r = book
	r.ReadBook()

	var w Writer
	// Why type assert suceess?
	// Because Interface `Reader` and `Writer` point to same concreate type (Book)
	w = r.(Writer)
	w.WriteBook()
}
