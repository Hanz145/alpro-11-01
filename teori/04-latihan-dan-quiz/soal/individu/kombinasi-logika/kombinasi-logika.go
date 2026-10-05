package main

import "fmt"

func main() {
	var p,q int
	fmt.Println("====== Kombinasi Logika ======")
	fmt.Print("Masukkan bilangan pertama : ")
	fmt.Scan(&p)
	fmt.Print("Masukkan bilangan kedua : ")
	fmt.Scan(&q)
	fmt.Print("Hasil : ")
	fmt.Print((p % 2 == 0) || (q % 2 == 0),  (p % 2 != 0) && (q % 2 != 0),  !(p == q))
}
