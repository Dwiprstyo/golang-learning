package main

import "fmt"

func main() {
	const (
		firtName string = "Dwi"
		lastName string = "Prasetyo"
	)
	fmt.Println(firtName)
	fmt.Println(lastName)
	// error
	// firtName = "Tidak bisa diubah"
	// lastName = "Tidak bisa diubah"
}
