//go:build ignore
// +build ignore

package main

import (
	"fmt"

	// anonymous import
	// There is no actual import, but the init function of the package will be called
	_ "tech.odes/explorer-go/syntax/function/lib"
)

// define init function
func init() {
	fmt.Println("invoke main init function ...")
}

func main() {
	// invoke main
	fmt.Println("invoke main function ...")
}
