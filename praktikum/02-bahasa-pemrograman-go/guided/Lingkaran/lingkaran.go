package main

import "fmt"

func main() {
	var pi float32 = 3.14
	var r float32
	fmt.Println("========= Lingkaran =========")
	fmt.Print("Masukan jari jari lingkaran: ")
	fmt.Scan(&r)
	luas := pi * r * r
	fmt.Println("========= Luas Lingkaran =========")
	fmt.Println("Luas = ", pi, " * ", r, " * ", r, " = ", luas)
}
