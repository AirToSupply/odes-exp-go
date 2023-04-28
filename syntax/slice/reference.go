//go:build ignore
// +build ignore

package main

import "fmt"

// [TIPS] transfer by reference
func update(slice []int) {
	slice[0] = 1000
}

func main() {
	slice := []int{1, 2, 3, 4}

	fmt.Printf("slice value: %v\n", slice)

	update(slice)

	fmt.Printf("slice value: %v\n", slice)
}
