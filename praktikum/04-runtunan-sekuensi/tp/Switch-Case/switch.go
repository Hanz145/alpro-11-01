package main

import "fmt"

func main() {
	fmt.Println("====== Switch Case ======")
	fmt.Println("Masukan bilangan dan lihat apakah bilangan tersebut > 10, <= 25, atau == 10 : ")
	fmt.Print("Masukkan bilangan: ")
	var bilangan int
	fmt.Scan(&bilangan)

	switch {
	case bilangan > 10:
		fmt.Println("Bilangan lebih besar dari 10")
	case bilangan > 25:
		fmt.Println("Bilangan lebih kecil atau sama dengan 25")
	case bilangan == 10:
		fmt.Println("Bilangan sama dengan 10")
	default:
		fmt.Println("Bilangan tidak memenuhi kondisi")
	}
}