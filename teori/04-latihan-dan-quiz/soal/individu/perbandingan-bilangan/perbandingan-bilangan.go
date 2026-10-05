package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("====== Perbandingan Bilangan ======")
	fmt.Print("Masukkan bilangan pertama: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan bilangan kedua: ")
	fmt.Scan(&b)
	fmt.Print(a>b)
	fmt.Print(a==b)
	fmt.Print(a<b)
}
