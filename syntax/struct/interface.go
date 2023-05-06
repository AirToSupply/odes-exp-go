//go:build ignore
// +build ignore

package main

import "fmt"

// interface is equivalent to `pointer`
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat Sleep ...")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog Sleep ...")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

// Implementation of polymorphism [OOP]
func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("Color :", animal.GetColor())
	fmt.Println("Kind :", animal.GetType())
}

func main() {
	// define interface
	var animal AnimalIF

	animal = &Cat{color: "Green"}
	animal.Sleep()

	animal = &Dog{color: "Yellow"}
	animal.Sleep()

	// Implementation of polymorphism
	fmt.Println("Implementation of polymorphism")
	showAnimal(&Cat{color: "Green"})
	showAnimal(&Dog{color: "Yellow"})
}
