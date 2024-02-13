package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	//---CARA CEPAT---
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	//--FUNC PAKE LEN--
	names := []string{"A", "B", "C"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	//--FUNC PAKE RANGE--
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
	//tanpa index
	for _, name := range names {
		fmt.Println(name)
	}
}
