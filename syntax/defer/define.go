//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	defer fmt.Println("Run defer block 1 ...")
	defer fmt.Println("Run defer block 2 ...")

	fmt.Println("Run main 1 ...")
	fmt.Println("Run main 2 ...")
}
