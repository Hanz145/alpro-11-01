package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Aku adalah pesulap yang bisa menukarkan nilai kedua variabel. Silahkan masukan angka nya Tuan/nyonya")
	fmt.Print("Variabel A: ")
	fmt.Scan(&a)
	fmt.Print("Variabel B: ")
	fmt.Scan(&b)
	c = a
	a = b
	b = c
	fmt.Println("Ini hasil nya tuan")
	fmt.Println("Variabel A :", a)
	fmt.Println("Variabel B :", b)
}
