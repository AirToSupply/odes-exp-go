//go:build ignore
// +build ignore

package main

import "fmt"

// iota refer to 0, and auto increment
const (
	CategoryBooks    = iota // 0
	CategoryHealth          // 1
	CategoryClothing        // 2
)

// use to operate
const (
	IgEggs         int = 1 << iota // 1 << 0 which is 00000001
	IgChocolate                    // 1 << 1 which is 00000010
	IgNuts                         // 1 << 2 which is 00000100
	IgStrawberries                 // 1 << 3 which is 00001000
	IgShellfish                    // 1 << 4 which is 00010000
)

// use to operates
const (
	_          = iota             // ignore first value by assigning to blank identifier
	KB float64 = 1 << (10 * iota) // 1 << (10*1)
	MB                            // 1 << (10*2)
	GB                            // 1 << (10*3)
	TB                            // 1 << (10*4)
	PB                            // 1 << (10*5)
	EB                            // 1 << (10*6)
	ZB                            // 1 << (10*7)
	YB                            // 1 << (10*8)
)

// auto increment by rows ?
// [TIPS] iota is only use to const block
const (
	Apple, Banana     = iota + 7, iota + 15 // 7, 15
	Cherimoya, Durian                       // 8, 16
	Elderberry, Fig                         // 9, 17
)

func main() {
	fmt.Printf("CategoryBooks=%v, CategoryHealth=%v, CategoryClothing=%v\n", CategoryBooks, CategoryHealth, CategoryClothing)

	fmt.Printf("IgEggs = %v, IgNuts = %v\n", IgEggs, IgNuts)

	fmt.Printf("KB = %v, GB = %v\n", KB, GB)

	fmt.Printf("Apple = %v\n", Apple)
	fmt.Printf("Banana = %v\n", Banana)
	fmt.Printf("Cherimoya = %v\n", Cherimoya)
	fmt.Printf("Durian = %v\n", Durian)
	fmt.Printf("Elderberry = %v\n", Elderberry)
	fmt.Printf("Fig = %v\n", Fig)
}
