package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func sayHelloTo(firstName string, lastName string)  {
	fmt.Println("Hello", firstName, lastName)
}

func getHello(name string) string  {
	return "Hello " + name
}

func getFullName() (string, string)  {
	return "Husnul", "Fikri"
}

func getCompleteName() (firstName, middleName, lastName string)  {
	firstName = "Husnul"
	middleName = "Fikri"
	lastName = "Husnulfk"

	return firstName, middleName, lastName
}

func main() {
	sayHello()
	sayHelloTo("Husnul", "FIkri")

	result := getHello("Husnul")
	fmt.Println(result)

	firstName, _ := getFullName()
	fmt.Println(firstName)

	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}