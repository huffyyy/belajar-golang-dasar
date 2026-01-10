package main

import "fmt"

func main() {
	name := "Husnullllllllll"

	if name == "Husnul" {
		fmt.Println("Hello Husnul")
	} else if name == "Fikri" {
		fmt.Println("Hello Fikri")
	} else {
		fmt.Println("Hi boleh kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}