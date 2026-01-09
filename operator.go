package main

import "fmt"

func main() {
	var a = 10
	var b = 15
	var c = a + b
	var d = 2
	var e = c * d
	
	fmt.Println(c)
	fmt.Println(e)

	var i = 10
	i += 10
	fmt.Println(i)

	i += 5
	fmt.Println(i)

	var j = 1
	j++
	fmt.Println(j)

	j--
	fmt.Println(j)

}