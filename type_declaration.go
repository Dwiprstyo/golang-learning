package main

import "fmt"

func main() {

	type NoKTP string

	var ktpTyo NoKTP = "111111111"

	var contoh string = "222222222"

	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpTyo)
	fmt.Println(contohKtp)
}
