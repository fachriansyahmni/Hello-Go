package main

import "fmt"

func main() {
	var benda [2]string

	benda[0] = "gunting"
	benda[1] = "jendela"

	fmt.Println(benda)

	var values = [2]int{
		190,
		2,
	}

	fmt.Println(values)
	fmt.Println(len(values))

}
