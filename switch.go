package main

import "fmt"

func main() {
	name := "Husnul"

	switch name {
	case "Husnul":
		fmt.Println("Hello Husnul")
	case "Fikri":
		fmt.Println("Hello Fikri")	
	default :
		fmt.Println("Hello boleh kenalan?")
	} 

	switch length := len(name); length > 5 {
	case true : 
	fmt.Print("Nama terlalu panjang")
	case false : 
	fmt.Println("Nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10 :
	fmt.Println("Nama terlalu panjang")
	case length > 5 : 
	fmt.Println("Nama cukup panjang")
	default : 
	fmt.Println("Nama sudah benar?")
	}
	
}