//go:build ignore
// +build ignore

package main

import "fmt"

// [TIPS] transfer by value
func update(arr [4]int) {
	arr[0] = 100
}

func main() {
	arr := [4]int{1, 2, 3, 4}

	fmt.Printf("Array value: %v\n", arr)

	update(arr)

	fmt.Printf("Array value: %v\n", arr)
}
