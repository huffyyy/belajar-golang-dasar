package main

import "fmt"

func main() {
	person := map[string]string{
		"name":   "Husnul",
		"address": "Jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["Title"] = "Buku Golang"
	book["Author"] = "Husnul Fikri"
	book["wrong"] = "ups"

	delete(book, "wrong")

	fmt.Println(book)
	fmt.Println(len(book))
}