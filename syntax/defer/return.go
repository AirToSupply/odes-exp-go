//go:build ignore
// +build ignore

package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer function invoke ...")
	return 0
}

func returnFunc() int {
	fmt.Println("return function invoke ...")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

/*
 * Order by return and defer
 */
func main() {
	returnAndDefer()
	// [Result]
	// return function invoke ...
	// defer function invoke ...
}
