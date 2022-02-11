package main

import "fmt"

func sumAll(numbers ...int) int{
	total := 0
	for _, number := range numbers{
		total += number
	}
	return total
}

func main(){
	total := sumAll(11, 10, 30, 50)
	fmt.Println(total)
}