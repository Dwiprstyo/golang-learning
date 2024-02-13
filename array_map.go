package main

import "fmt"

func main() {
	//len(map) jumlah data
	//map[key] get data di map
	//map[key] = value mengubah data di map
	//make(map[TypeKey]TypeValue) membuat map baru
	//delete(map, key) // menghapus data di map dengan key

	//-----------cara lama--------------
	// var person map[string]string = map[string]string{}
	// person["name"] = "Dwi"
	// person["address"] = "Bandung"

	//cara cepat
	person := map[string]string{
		"name":    "Dwi",
		"address": "Bandung",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Dwi"
	book["ups"] = "Salah"

	fmt.Println(book, "book")

	delete(book, "ups")
	fmt.Println(book, "book didelete")
}
