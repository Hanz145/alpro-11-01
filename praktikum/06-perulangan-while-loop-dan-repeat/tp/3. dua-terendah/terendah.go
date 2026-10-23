package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var namaTerendah, namaKeduaTerendah string
	var nilaiTerendah, nilaiKeduaTerendah int

	nilaiTerendah = 100
	nilaiKeduaTerendah = 100

	var i int = 0
	for i < n {
		var nama string
		var nilai int
		fmt.Scan(&nama, &nilai)

		if nilai < nilaiTerendah {
			nilaiKeduaTerendah = nilaiTerendah
			namaKeduaTerendah = namaTerendah
			nilaiTerendah = nilai
			namaTerendah = nama
		} else if nilai < nilaiKeduaTerendah {
			nilaiKeduaTerendah = nilai
			namaKeduaTerendah = nama
		}

		i++
	}
	fmt.Println(namaTerendah)
	fmt.Println(namaKeduaTerendah)
}