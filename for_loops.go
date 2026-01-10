package main

import (
	"fmt"
)

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	names := []string{"Husnul", "Fikri"}
	for index, name:= range names {
		fmt.Println("index", index, "=", name)
	}
}