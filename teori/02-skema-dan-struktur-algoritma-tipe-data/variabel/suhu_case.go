package main

import "fmt"

func main() {
	var umur int8 = 10
	var suhu float32 = 36.4
	fmt.Println("Umur: ", umur)
	fmt.Println("Suhu: ", suhu)

	fmt.Println("Almat memori var suhu:", &suhu)
	fmt.Println("Almat memori var umut:", &umur)
}
