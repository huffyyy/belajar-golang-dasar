package main

import (
	"fmt"
	"strconv"
	"strings"
)

func horizontalBintang(n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += "*"
	}
	return result
}

func kotakBintang(n int) string {
	result := ""

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result += "*"
		}
		result += "\n"
	}
	return  result
}

func kotakAngka(n int) string {
	var b strings.Builder
	angka := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			b.WriteString(strconv.Itoa(angka))
			b.WriteString(" ")
			angka++
		}
		b.WriteString("\n")
	}
	return b.String()
}

func main() {
	var n int
	fmt.Print("Masukkan n: ")
	fmt.Scanln(&n)

	fmt.Println(horizontalBintang(n))
	fmt.Println(kotakBintang(n))
	fmt.Println(kotakAngka(n))
}
