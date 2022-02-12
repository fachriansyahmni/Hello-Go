package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "UPS!"
	}
}

func main() {
	var cek interface{} = Ups(1)
	fmt.Println(cek)
}
