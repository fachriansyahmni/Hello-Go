package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	Rully := Customer{"Rully", "Jl.Merdeka", 20}
	Rully.sayHello("Zio")
}
