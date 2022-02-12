package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var eko Customer
	eko.Name = "Eko Setiawan"
	eko.Address = "Jl. Setiabudi no 12"
	eko.Age = 21

	fmt.Println(eko)

	junaedi := Customer{
		Name:    "Junaedi",
		Address: "Jl. Buah batu",
		Age:     20,
	}

	fmt.Println(junaedi)

	budi := Customer{"Budi", "Jl. Kacapi", 15}
	fmt.Println(budi)

}
