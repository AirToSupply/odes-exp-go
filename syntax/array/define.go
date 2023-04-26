//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	// Methon 1: fixed length and do not assige value
	var arr1 [10]int
	// value: [0 0 0 0 0 0 0 0 0 0]
	fmt.Printf("arr1 type: %T, arr1 value: %v\n", arr1, arr1)

	// Methon 2: fixed length and partial assige value
	arr2 := [10]int{1, 2, 3, 4}
	// value: [1 2 3 4 0 0 0 0 0 0]
	fmt.Printf("arr1 type: %T, arr1 value: %v\n", arr2, arr2)
}
