package main

import "fmt"

func main() {
	name := "Jokiii"

	if name == "Dwi" {
		fmt.Println("Hello", name)
	} else if name == "Joki" {
		fmt.Println("Hello Joki")
	} else {
		fmt.Println("Hi, Kenalan Dong")
	}

	// cara cepat
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
