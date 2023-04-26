//go:build ignore
// +build ignore

package main

import (
	"fmt"

	"tech.odes/explorer-go/syntax/function/lib"
)

// define init function
func init() {
	fmt.Println("invoke main init function ...")
}

func main() {
	// invoke Run function by lib package
	lib.Run()
	// invoke main
	fmt.Println("invoke main function ...")
}
