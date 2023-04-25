//go:build ignore
// +build ignore

package main

import "fmt"

// define muti variables outside function body
// [ERROR] syntax error: non-declaration statement outside function body
// [TIPS] Do not automatic derivation outside function body
// k, v := "key", "value"
var k, v = "key", "value"

// define muti variables outside function body by muti-line
var (
	id   = 1
	name = "hudi"
	age  = 17
)

func main() {

	// define muti variables in function body
	x, y := 3, 4
	fmt.Printf("x = %v, y = %v\n", x, y)

	// define muti variables and ignore partial variables
	pointX, _, pointY, _ := 1, "Unknow", 2, "Unknow"
	fmt.Printf("pointX = %v, pointY = %v\n", pointX, pointY)

	// define muti variables without assigning a value
	var p, q int
	fmt.Printf("p = %v, q = %v\n", p, q) //  p = 0, q = 0

	fmt.Printf("k = %v, v = %v\n", k, v)

	var info = `
	[User Info]
      [Id]    %v
	  [Name]  %v
	  [Age]   %v
	`
	fmt.Printf(info, id, name, age)
}
