//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	// define and do not allocate space
	var m map[string]string
	if m == nil {
		fmt.Println("map is nil")
	}

	// Method 1: allocate space of map and specify map size by make function before using
	m = make(map[string]string, 10)
	m["one"] = "java"
	m["two"] = "c++"
	m["three"] = "python"
	fmt.Println("m =", m)

	// Method 2:
	m1 := make(map[int]string)
	m1[1] = "java"
	m1[2] = "c++"
	m1[3] = "python"
	fmt.Println("m1 =", m1)

	// Method 3: directed define
	m2 := map[string]string{
		"one":   "php",
		"two":   "golang",
		"three": "java",
	}
	fmt.Println("m2 =", m2)
}
