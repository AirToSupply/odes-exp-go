//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	s1 := s[0:2]
	// index from 0 to 2 - 1
	fmt.Println("s1 =", s1) // [1 2]

	// start index is not specfiy, it is 0
	fmt.Println("s[:2] =", s[:2]) // [1 2]

	// end index is not specfiy, it is length of slice - 1
	fmt.Println("s[1:] =", s[1:]) // [2 3]

	// start index and end index are not specfied
	fmt.Println("s[:] =", s[:]) // [1 2 3]

	// [Qos] When change value of s1[0], Has the value of s[0] changed?
	s1[0] = 100
	fmt.Println("s =", s)
	fmt.Println("s1 =", s1)

}
