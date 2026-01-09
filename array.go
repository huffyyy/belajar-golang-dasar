package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "Husnul"
	names[1] = "Fikri"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [...]int {
		1,
		2,
		3,
	}
	fmt.Println(values)	
	fmt.Println(len(values))	
	fmt.Println(values[2])	
	values[0] = 10
	fmt.Println(values)	
}