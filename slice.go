package main

import "fmt"

func main() {
	names := [...]string{"Java", "PHP", "C", "Ruby", "Go", "Pyhton"}
	sliceLow := names[2:]
	sliceHigh := names[:2]
	sliceLowHigh := names[4:6]

	fmt.Println(sliceLow)
	fmt.Println(sliceHigh)
	fmt.Println(sliceLowHigh[0])

	days :=[...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daysSlice1 := days[5:]
	daysSlice1[0] = "sabtu baru"
	daysSlice1[1] = "minggu baru"
	fmt.Println(days)

	daysSlice2 := append(daysSlice1, "libur baru")
	daysSlice2[0] = "ups"	
	fmt.Println(daysSlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Husnul"
	newSlice[1] = "Husnul"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	formSlice := days[:]
	toSlice := make([]string, len(formSlice), cap(formSlice))

	copy(toSlice, formSlice)

	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}