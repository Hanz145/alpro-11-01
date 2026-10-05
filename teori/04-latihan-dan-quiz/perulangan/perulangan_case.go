package main

import "fmt"

func main() {
	var jumlahbilangan, bilangan, baris1, baris4, hasil, a int
	fmt.Scan(&jumlahbilangan)
	for a = 0; a < jumlahbilangan; a++ {
		fmt.Scan(&bilangan)
		baris1 = bilangan / 1000
		baris4 = bilangan % 10
		hasil = hasil + baris1 + baris4
	}
	println(hasil)
}