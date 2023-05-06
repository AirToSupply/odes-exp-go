//go:build ignore
// +build ignore

package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human eat ...")
}

func (this *Human) Walk() {
	fmt.Println("Human walk ...")
}

type SuperMan struct {
	// => SuperMan extends Human
	Human
	level int
}

// => override
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan eat ...")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan fly ...")
}

func (this *SuperMan) Print() {
	fmt.Println("name :", this.name)
	fmt.Println("sex :", this.sex)
	fmt.Println("level :", this.level)
}

func main() {
	human := Human{name: "Zh", sex: "F"}
	human.Eat()
	human.Walk()

	// Method 1
	superMan := SuperMan{human, 88}
	// Method 2
	superMan.name = "superman"
	superMan.sex = "M"
	superMan.level = 1
	superMan.Eat()
	superMan.Walk() // invoke parent class method
	superMan.Fly()
	superMan.Print()
}
