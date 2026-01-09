package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 75
	var lulusAbsensi bool = absensi > 70

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)
}