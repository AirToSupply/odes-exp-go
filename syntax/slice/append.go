//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	numbers := make([]int, 3, 5)
	fmt.Printf("numbers type: %T, numbers value: %v, len: %d, cap: %d\n", numbers, numbers, len(numbers), cap(numbers))

	// append element
	// len from 3 to 4, cap is 5
	numbers = append(numbers, 1)
	fmt.Printf("numbers type: %T, numbers value: %v, len: %d, cap: %d\n", numbers, numbers, len(numbers), cap(numbers))

	// append element
	// len from 4 to 5 , cap is 5
	numbers = append(numbers, 2)
	fmt.Printf("numbers type: %T, numbers value: %v, len: %d, cap: %d\n", numbers, numbers, len(numbers), cap(numbers))

	// append element when slice is filled (len > cap)
	// len from 5 to 6 , cap from 5 to 10
	// It indicates that slice capacity has been expanded to twice the original capacity
	numbers = append(numbers, 3)
	fmt.Printf("numbers type: %T, numbers value: %v, len: %d, cap: %d\n", numbers, numbers, len(numbers), cap(numbers))
}
