//go:build ignore
// +build ignore

package main

import "fmt"

// [TIPS] transfer by reference
func update(city map[string]string) {
	city["England"] = "London"
}

func main() {
	city := map[string]string{
		"China": "Beijing",
		"Japan": "Tokyo",
		"USA":   "NewYork",
	}

	fmt.Println("update before city =", city)

	update(city)

	fmt.Println("update after city =", city)
}
