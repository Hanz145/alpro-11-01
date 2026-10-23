package main

import "fmt"

func main() {
	var nilai int
	var a string
	fmt.Print("Masukan Nilai: ")
	fmt.Scan(&nilai)

	if nilai < 50 && nilai > -1 {
		a = "E"
	} else if nilai >= 50 && nilai < 70 {
		a = "C"
	} else if nilai >= 70 && nilai < 85 {
		a = "B"
	} else if nilai >= 85 && nilai <= 100 {
		a = "A"
	}

	fmt.Print("Siswa Mendapatkan Nilai ", a)
}
