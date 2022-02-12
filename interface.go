package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Food struct {
	Name string
}

func (food Food) GetName() string {
	return food.Name
}

func main() {
	var eko Person
	eko.Name = "Eko Setiawan"
	sayHello(eko)

	food := Food{"Manggo"}
	sayHello(food)
}
