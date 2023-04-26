//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4}

	// Method 1: iterator by length
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	fmt.Println("-----------------")

	// Method 2: iterator by range
	for index, value := range arr {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}
}
