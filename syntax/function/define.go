//go:build ignore
// +build ignore

package main

import "fmt"

// Define a function with a return value
func add(x int, y int) int {
	return x + y
}

// Define a function with a return muti values
func swap(x, y string) (string, string) {
	return y, x
}

// Give a name to the return value
func shift(x, y, delta int) (pointX int, pointY int) {
	pointX = x + delta
	pointY = y + delta
	return pointX, pointY
}

func main() {
	// add
	x, y := 4, 5
	fmt.Printf("x + y = %v\n", add(x, y))

	// swap
	a, b := "left", "right"
	a, b = swap(a, b)
	fmt.Printf("swap: a=%v, b=%v\n", a, b)

	// shift
	delta := 10
	x, y = shift(x, y, delta)
	fmt.Printf("shift point: x=%v, y=%v\n", x, y)
}
