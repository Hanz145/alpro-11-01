package main

import "fmt"

func main() {
	//varibel
	var NamaSiswa string
	var NilaiBahasaInggris, NilaiMatematika int
	//input
	fmt.Print("Nama: ")
	fmt.Scan(&NamaSiswa)

	fmt.Print("Nilai Bahasa Inggris: ")
	fmt.Scan(&NilaiBahasaInggris)

	fmt.Print("Nilai Matematika: ")
	fmt.Scan(&NilaiMatematika)
	//rata rata
	jumlah := NilaiBahasaInggris + NilaiMatematika
	ratarata := jumlah / 2
	//keluaran
	fmt.Print("Rata-Rata: ", ratarata)
}
