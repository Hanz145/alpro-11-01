package main

import "fmt"
import "bufio"
import "os"

func main() {
	var nama string
	var nilai float32
	fmt.Println("====== Program Penilaian ======")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan nama: ")
	scanner.Scan()
	nama = scanner.Text()
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	fmt.Println("====== Hasil Penilaian ======")
	fmt.Println("Nama:", nama)
	fmt.Println("Nilai:", nilai)

	if nilai >= 90 && nilai <= 100 {
		fmt.Println("Nilai A")
	}
	if nilai >= 80 && nilai < 90 {
		fmt.Println("Nilai B")
	}
	if nilai >= 70 && nilai < 80 {
		fmt.Println("Nilai C")
	}
	if nilai >= 60 && nilai < 70 {
		fmt.Println("Nilai D")
	}
	if nilai < 60 {
		fmt.Println("Nilai F")
	}
}
