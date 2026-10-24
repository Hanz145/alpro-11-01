package main


import "fmt"

func main(){
	var nilai int // var nilai

	for {
		fmt.Print("Masukan Nilai: ")
		fmt.Scan(&nilai)
		if nilai >= 60 && nilai < 101{
			fmt.Println("Anda lulus ujian")
		} else if nilai < 60 && nilai > -1{
			fmt.Println("Anda tidak lulus ujian")
		} else if nilai == -1 {
			break
		} else {
			fmt.Println("tidak ada nilai, mohon masukan angka yang valid (0 - 100)")
		}
	}
}