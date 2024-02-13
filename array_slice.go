package main

import "fmt"

func main() {
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray, "iniArray")
	fmt.Println(iniSlice, "iniSlice")

	//SLICE
	names := [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"}

	slice1 := names[4:6]
	fmt.Println(slice1, "4:6 ambil lenfh 4-6")

	slice2 := names[:3]
	fmt.Println(slice2, ":3 ambil 3 diawal")

	slice3 := names[3:]
	fmt.Println(slice3, "3: ambil semua dimulai dari legth 3")

	slice4 := names[:]
	fmt.Println(slice4, ":  ambil semua")

	var slice5 []string = names[:]
	fmt.Println(slice5, ":  ambil semua pake variable")

	//APPEND
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println(days, "days-1")

	daySlice1 := days[5:] // Sabtu, Minggu
	fmt.Println(daySlice1, "daySlice1-1")

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1, "daySlice1-2")
	fmt.Println(days, "days-2")

	daySlice2 := append(daySlice1, "Libur Baru")
	fmt.Println(daySlice2, "daySlice2")

	//MAKE
	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Dwi"
	newSlice[1] = "Dwi"

	fmt.Println(newSlice, "newSlice")
	fmt.Println(len(newSlice), "LEN newSlice")
	fmt.Println(cap(newSlice), "CAP, newSlice")

	newSlice2 := append(newSlice, "Dwi")
	fmt.Println(newSlice2, "newSlice2")
	fmt.Println(len(newSlice2), "LEN newSlice2")
	fmt.Println(cap(newSlice2), "CAP, newSlice2")

	newSlice2[0] = "Budi"
	fmt.Println(newSlice)
	fmt.Println(newSlice2)

	//COPY SLICE
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice, "fromSlice")
	fmt.Println(toSlice, "toSlice")
}
