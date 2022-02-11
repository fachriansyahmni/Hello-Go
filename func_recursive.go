package main

import "fmt"

func factorialLoop(value int) int {
	rslt := 1
	for i := value; i > 0; i-- {
		rslt *= i
	}
	return rslt
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	loop := factorialLoop(5)
	fmt.Println(loop)
	fmt.Println(factorialRecursive(5))
}
