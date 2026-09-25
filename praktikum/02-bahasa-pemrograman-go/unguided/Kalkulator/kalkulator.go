package main

import "fmt"

func main() {
	var a, b int
	for i := 1; i < 2; i = i + 0 {
		fmt.Println("========= Masukan Dua Angka =========")
		fmt.Print("Bilangan A: ")
		fmt.Scan(&a)
		fmt.Print("Bilangan B: ")
		fmt.Scan(&b)
		if b == 0 {
			fmt.Println("maaf, B tidak bisa berupa 0")
		} else {
			break
		}
	}
	tambah := a + b
	kurang := a - b
	kali := a * b
	bagi := a / b
	persen := a % b
	fmt.Println("========= Hasil Bilangan =========")
	fmt.Println("a + b = ", tambah)
	fmt.Println("a - b = ", kurang)
	fmt.Println("a * b = ", kali)
	fmt.Println("a / b = ", bagi)
	fmt.Println("a % b = ", persen)
}
