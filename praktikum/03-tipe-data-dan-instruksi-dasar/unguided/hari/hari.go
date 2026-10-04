package main

import "fmt"

func main() {
	var tahun, bulan, minggu, hari, sisa int32
	fmt.Println("====== Banyak Hari ======")
	fmt.Print("Hari : ")
	fmt.Scan(&hari)
	tahun = hari / 365
	sisa = hari % 365
	bulan = sisa / 30
	sisa = sisa % 30
	minggu = sisa / 4
	sisa = sisa % 4
	fmt.Println("====== Konverenis ke Tahun, Bulan, Minggu ======")
	fmt.Println("Tahun : ", tahun)
	fmt.Println("Bulan : ", bulan)
	fmt.Println("Minggu : ", minggu)
	fmt.Println("sisa : ", sisa)
}
