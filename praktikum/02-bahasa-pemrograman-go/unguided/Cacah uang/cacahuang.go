package main

import "fmt"

func main() {
	var jumlahuang int32
	fmt.Println("=========Jumlah Uang=========")
	fmt.Print("Masukan Jumlah Uang: ")
	fmt.Scan(&jumlahuang)
	uang10k := jumlahuang / 10000
	sisauang := jumlahuang % 10000
	uang5k := sisauang / 5000
	sisauang = sisauang % 5000
	uang1k := sisauang / 1000
	sisauang = sisauang % 1000
	fmt.Println("=========Pecahan Uang=========")
	fmt.Println("Uang 10k: ", uang10k)
	fmt.Println("Uang 5k: ", uang5k)
	fmt.Println("Uang 1k: ", uang1k)
	fmt.Println("Uang sisa: ", sisauang)
	fmt.Println("Jumlah uang ", jumlahuang)
}
