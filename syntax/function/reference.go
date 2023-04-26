//go:build ignore
// +build ignore

package main

import "fmt"

/*
 * transfer by value VS. transfer by reference
 */
func main() {
	var a int = 100
	var b int = 200

	fmt.Printf("swap by value before a=%d, b=%d\n", a, b)
	swapByValue(a, b)
	fmt.Printf("swap by value after a=%d, b=%d\n", a, b)

	fmt.Printf("swap by ref before a=%d, b=%d\n", a, b)
	swapByRef(&a, &b)
	fmt.Printf("swap by ref after a=%d, b=%d\n", a, b)
}

// [TIPS] transfer by value
func swapByValue(x, y int) int {
	var temp int
	temp = x
	x = y
	y = temp
	return temp
}

// [TIPS] transfer by reference
func swapByRef(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
