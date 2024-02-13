package main

import "fmt"

func main() {

	//--LOOP BREAK--
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan Break", i)
	}
}
