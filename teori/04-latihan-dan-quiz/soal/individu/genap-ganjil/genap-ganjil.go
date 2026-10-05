package main

import "fmt"

func main() {
	var n int
	fmt.Println("====== Genap Ganjil or Genap ======")
	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&n)
	fmt.Print("keterangan\ngenap = true\nganjil = false")
	fmt.Print("\nHasil : ",n % 2 == 0)
}
