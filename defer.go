package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApp() {
	defer logging()
	fmt.Println("Start APP")
}

func runApp2(value int) {
	defer logging()
	fmt.Println("Start APP")
	bagi := 10 / value
	fmt.Println(bagi)
}

func main() {
	// runApp()
	runApp2(0)
}
