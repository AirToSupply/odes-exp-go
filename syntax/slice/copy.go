//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	s := make([]int, 4)

	fmt.Println("s =", s) // [0, 0, 0, 0]

	numbers := []int{1, 2, 3}

	// Copy the slices of the underlying array together by `copy` function
	// copy slice `numbers` to slice `s`
	copy(s, numbers)
	fmt.Println("s =", s) // [1 2 3 0]

	// [Qos] When change value of s[1], Has the value of numbers[1] changed?
	s[1] = 100
	fmt.Println("s =", s)
	fmt.Println("numbers =", numbers)

}
