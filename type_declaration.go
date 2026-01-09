package main

import "fmt"

func main() {
	type noKTP string

	var ktpHusnul noKTP = "11111111"

	var contoh = "22222222"
	var contohKTP = noKTP(contoh)

	fmt.Println(ktpHusnul)
	fmt.Println(contoh)
	fmt.Println(contohKTP)
}