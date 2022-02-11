package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You're Blocked!")
	} else {
		fmt.Println("Welcome!")
	}
}

func main() {
	blackList := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blackList)
	registerUser("Eko", blackList)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
}
