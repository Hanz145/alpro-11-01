package main

import "fmt"

func main() {
	var x, y int
	//memasukan nilai
	fmt.Println("Masukan (jumlah anggota, keluarga dan jumlah kue) : ")
	fmt.Scan(&x)
	fmt.Scan(&y)
	//mencari sisa bagi
	sisahasil := y % x
	//sisa bagi
	fmt.Print("Sisa kue : ")
	fmt.Print(sisahasil)
}
