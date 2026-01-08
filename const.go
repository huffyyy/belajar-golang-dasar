package main

import "fmt"

func main() {
	const (
		firstName = "Husnul"
		lastName  = "Fikri"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// error
	// firstName = "Fikri"
	// lastName = "Husnul"
}