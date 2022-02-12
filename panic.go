package main

import "fmt"

func endApp() {
	msg := recover()
	if msg != nil {
		fmt.Println("Eror dengan pesan ", msg)
	}
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("App Error!!")
	}
	fmt.Println("App Running")
}

func main() {
	runApp(true)
}
