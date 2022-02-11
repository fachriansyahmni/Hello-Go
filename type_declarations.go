package main

import "fmt"

func main() {
	type NoKTP string

	var noKtpEko NoKTP = "12345678"
	fmt.Println(noKtpEko)
}
