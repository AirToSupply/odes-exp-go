package lib

import "fmt"

// define init function
func init() {
	fmt.Println("invoke tool init function by lib package...")
}

func Run() {
	fmt.Println("invoke tool Run function by lib package...")
}
