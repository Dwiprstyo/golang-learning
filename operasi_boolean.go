package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 81

	var luluNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = luluNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)
}
