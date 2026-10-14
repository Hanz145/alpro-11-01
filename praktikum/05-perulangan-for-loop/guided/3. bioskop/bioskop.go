package main

import "fmt"

func main(){
	var t int
	fmt.Print("input jumlah transaksi: ")
	fmt.Scan(&t)

	totalpendapatan := 0

	for i := 1; i <= t; i++{
		var jenisvilm, hari string
		var jumlahtiket int

		fmt.Print("input transaksi ke %d (jenis, hari, tiket): ",i)
		fmt.Scan(&jenisvilm, &hari, &jumlahtiket)
		
		hargasatuan := 0

		if jenisvilm == "Reguler"{
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" || {
				hargasatuan = 40000
			} else {
				hargasatuan = 50000
			}
		} else if jenisvilm == "3D" {
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" || {
				hargasatuan = 60000
			} else {
				hargasatuan = 70000
			}
		} else if jenisvilm == "IMAX" {
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" || {
				hargasatuan = 80000
			} else {
				hargasatuan = 90000
			}
		} else {
			hargasatuan = 90000
		}

		totalharga

	}
}