package main

import "fmt"

func main (){
	var n, hasil int // var n dan hasil
	var count int // var pendeteksi

	fmt.Println("====== Bilaangan Prima ======")
	i := 1
	fmt.Print("Bilangan: ")
	fmt.Scan(&n) //input
	limit := n

	fmt.Print("Faktor: ")
	for { // perulangan 1
		hasil = n % i
		if hasil == 0 { // jika n == 0
			fmt.Print(i," ")
			count += 1
		}
		i += 1
		if i > limit { // break
			break
		}
	}

	fmt.Print("\nPrima: ")
	if count < 3 {
		fmt.Print("true") // jika prima maka hanya mengitung dua angka
	} else {
		fmt.Print("False") // jika tidak dia akan menghitung lebih dari dua
	}
}