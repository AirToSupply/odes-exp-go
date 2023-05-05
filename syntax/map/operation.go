//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	// define
	m := map[string]string{
		"China": "Beijing",
		"Japan": "Tokyo",
		"USA":   "NewYork",
	}

	// append
	m["South Korea"] = "Seoul"

	// select
	for key, value := range m {
		fmt.Println("key =", key, ", value =", value)
	}

	// delete
	delete(m, "South Korea")

	// modify
	m["USA"] = "DC"

	for key, value := range m {
		fmt.Println("key =", key, ", value =", value)
	}
}
