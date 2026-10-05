package main

import "fmt"

func main() {
	var a, hasil int
	for loop := 0; loop < 1; loop = loop + 0 {
		fmt.Println("====== Bilangan Awal ======")
		fmt.Print("Masukkan jumlah bilangan: ")
		fmt.Scan(&a)
		if a > 99 {
			fmt.Println("Masukkan bilangan dua digit saja")
		} else {
			break
		}
	}
	bil1 := a / 10
	bil2 := a % 10
	hasil = bil1*1000 + bil1*100 + bil2*10 + bil2
	fmt.Println("====== Hasil ======")
	fmt.Print(hasil)
}
