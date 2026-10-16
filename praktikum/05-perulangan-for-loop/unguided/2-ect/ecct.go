package main

import "fmt"

func main(){

	var siswa int // var mengatur banyak siwswa
	fmt.Println("====== ECCT Siswa ======")
	fmt.Print("Jumlah Siswa: ")
	fmt.Scan(&siswa) // input
	loop := 0 // var loop
	var  P, V, GA, FS, CI, TC, ratarata float32 //untuk nilai
	count := 1
	countlulus := 0 // menghitung siswa lulus

	for { // loop
		fmt.Println("====== Nilai Siswa ======")
		fmt.Println("Siswa: ", count)
		fmt.Print("Nilai P: ")
		fmt.Scan(&P)

		fmt.Print("Nilai V: ")
		fmt.Scan(&V)

		fmt.Print("Nilai GA: ")
		fmt.Scan(&GA)

		fmt.Print("Nilai FS: ")
		fmt.Scan(&FS)

		fmt.Print("Nilai CI: ")
		fmt.Scan(&CI)

		fmt.Print("Nilai TC: ")
		fmt.Scan(&TC)

		ratarata = (P + V + GA + FS + CI + TC)/6

		fmt.Println("Nilai Minimal Rata-Rata: 3,33")
		fmt.Println("Rata-Rata: ", ratarata)

		if ratarata >= 3.33 {
			fmt.Println("Status: Ya") // jika di atas rata rata
			countlulus += 1
		} else {
			fmt.Println("Lulus: Tidak") // jika tidak
		}
		
		loop += 1
		if loop == siswa{
			break
		}
	}
	fmt.Print("Sisiwa Yang Lulus: ",countlulus)
}
