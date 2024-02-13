package main

import "fmt"

func main() {
	name := "Dwiii"

	//default
	switch name {
	case "Dwi":
		fmt.Println("Hello Dwi")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Nama Tidak Ketemu")
	}

	//cara cepat
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama Sudah Pas")
	}

	//switch tanda kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu Panjang")
	case length > 5:
		fmt.Println("Nama lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}
