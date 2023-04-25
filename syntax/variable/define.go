//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	// Method 1: specify variable type without assigning a value
	var a int
	fmt.Printf("a = %d\n", a) // Print default values

	// Method 2: specify variable type and assige a value
	var b int = 10
	fmt.Printf("b = %d\n", b)

	// Method 3: not specify variable type and compiler automatic derivation
	var c = 3.141592653
	fmt.Printf("c = %v, type = %T\n", c, c) // float64

	// Method 4: ingnore key of `var`
	d := "variable defines"
	fmt.Printf("d = %v, type = %T\n", d, d) // string
}
