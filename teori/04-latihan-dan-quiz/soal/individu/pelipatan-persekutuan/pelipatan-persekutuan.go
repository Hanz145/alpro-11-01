package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Println("====== Pelipatan Persekutuan ======")
	fmt.Print("Masukkan bilangan pertama : ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan kedua : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan bilangan ketiga : ")
	fmt.Scan(&b)
	fmt.Print("Hasil : ")
	fmt.Print((n % a == 0) && (n % b == 0))
}
