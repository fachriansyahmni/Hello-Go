package main

import "fmt"

func main() {
	counter := 0
	nama := "jamal"
	increment := func() {
		nama = "budi"
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(nama)
}
