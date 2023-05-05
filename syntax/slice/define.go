//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	// Method 1: assige value
	slice := []int{1, 2, 3}
	// len = 3, cap = 3
	fmt.Printf("slice type: %T, slice value: %v, len: %d, cap: %d\n", slice, slice, len(slice), cap(slice))

	// Method 2: define and do not allocate space
	var s1 []int
	// s1 type: []int, s1 value: [], len: 0, cap: 0
	fmt.Printf("s1 type: %T, s1 value: %v, len: %d, cap: %d\n", s1, s1, len(s1), cap(s1))
	// if not allocate space, s1 value is nil
	if nil == s1 {
		fmt.Println("s1 is nil")
	}
	// allocate space and specify length size is 3
	s1 = make([]int, 3)
	s1[0] = 100
	fmt.Printf("s1 type: %T, s1 value: %v, len: %d, cap: %d\n", s1, s1, len(s1), cap(s1))

	// Method 3: ingnore key of `var`
	s2 := make([]int, 3)
	fmt.Printf("s2 type: %T, s2 value: %v, len: %d, cap: %d\n", s2, s2, len(s2), cap(s2))

	// Method 4: allocate space and specify length size is 3 and capacity size 5
	s3 := make([]int, 3, 5)
	fmt.Printf("s3 type: %T, s3 value: %v, len: %d, cap: %d\n", s3, s3, len(s3), cap(s3))
}
