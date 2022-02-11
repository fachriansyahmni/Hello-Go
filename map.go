package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Eko",
		"address": "Bandung",
	}

	fmt.Println(person)

	book := make(map[string]string)
	fmt.Println(book)

}
