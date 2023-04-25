//go:build ignore
// +build ignore

package main

import "fmt"

// Method 3: Use to enum
const (
	UNKNOW = 0
	FEMALE = 1
	MALE   = 2
)

func main() {
	// Method 1: specify variable type
	const LENGTH int = 10
	const WIDTH int = 10
	fmt.Printf("Area = %v\n", LENGTH*WIDTH)

	// Method 2: not specify variable type and compiler automatic derivation
	const x, y, z = 3, 4, 5
	fmt.Printf("Location = (%v, %v, %v)\n", x, y, z)

	fmt.Printf("[UNKNOW, FEMALE, MALE] = (%v, %v, %v)\n", UNKNOW, FEMALE, MALE)
}
