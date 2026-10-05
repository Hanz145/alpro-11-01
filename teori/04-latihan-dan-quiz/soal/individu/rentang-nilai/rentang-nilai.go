package main

import "fmt"

func main() {
	var x, low, high int
	fmt.Println("====== Rentang Nilai ======")
	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&x)
	fmt.Print("Masukkan batas bawah : ")
	fmt.Scan(&low)
	fmt.Print("Masukkan batas atas : ")
	fmt.Scan(&high)
	fmt.Print("Hasil : ")
	fmt.Print((x >= low) && (x <= high))
}
