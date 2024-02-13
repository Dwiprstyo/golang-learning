package main

import "fmt"

func main() {
	//1
	var names [3]string
	names[0] = "M"
	names[1] = "Dwi"
	names[2] = "Prasetyo"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//2
	var values = [...]int{1, 2, 3, 4, 100}

	fmt.Println(values)

	//3
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)
}
