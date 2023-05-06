//go:build ignore
// +build ignore

package main

import "fmt"

type Hero struct {
	// Capitalizing the first letter of an attribute is equivalent to an access modifier at the `public`` level
	Name string
	Ad   int
	// Attribute lowercase is equivalent to priavte level access modifier
	Level int
}

func (this *Hero) Show() {
	fmt.Println("Name: ", this.Name)
	fmt.Println("Ad: ", this.Ad)
	fmt.Println("Level: ", this.Level)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(name string) {
	this.Name = name
}

func main() {
	hero := Hero{
		Name:  "Zh",
		Ad:    100,
		Level: 1,
	}

	fmt.Println("[Print Hero Info]")
	hero.Show()

	hero.SetName("Foo")

	fmt.Println("[Update Hero Name]")
	hero.Show()
}
