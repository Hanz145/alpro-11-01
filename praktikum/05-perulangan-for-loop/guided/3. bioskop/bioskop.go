package main

import "fmt"

func main() {

	var hari, tiket string // var hari sama jenis tiket
	var pesanan, total, jmlhtiket int // var banyak nya pesanan dan total harga pesanan
	fmt.Println("Silahkan Pesan tiket bioskop yang kamu mau")
	fmt.Println("==========================")
	fmt.Println("Reguler = 40000")
	fmt.Println("3D = 60000")
	fmt.Println("IMAX = 80000")
	fmt.Println("Jumat - Minggu = +10000")
	fmt.Println("==========================")
	fmt.Println("diskon 10% jika membeli 5 tiket dalam sehari")
	fmt.Println("jika ada kesalahan dalam memasukan hari, maka hari = Senin")
	fmt.Println("jika ada kesalahan dalam nama tiket, maka tiket = Reguler")
	fmt.Println("jika ada kesalahan dalam memasukan jumlah tiket, maka tiket = 1")
	fmt.Print("Banyak Pesanan: ")
	fmt.Scan(&pesanan)// berapa banyak tiket yang akan di pesan
	loop := 0 //var pengatur loop
	fmt.Print("\n ")


	for { // loop utama
		fmt.Println("====== Tiket ======")
		fmt.Print("Jenis Tiket: ")
		fmt.Scan(&tiket)
		fmt.Print("Hari (diawali kapital): ")
		fmt.Scan(&hari)
		fmt.Print("Jumlah Tiket: ")
		fmt.Scan(&jmlhtiket)
		hargatiket := 0 //menampilkan harga sesuai jenis tiket

		//mengecek harga tiket
		if tiket == "IMAX"{
			hargatiket = 80000
		} else if tiket == "3D" {
			hargatiket = 60000
		} else {
			hargatiket = 40000
		}

		//mengecek hari
		if hari == "Jumat" || hari == "Sabtu" || hari == "Minggu" {
			hargatiket += 10000
		} else if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
			hargatiket += 0
		} else {
			hargatiket += 0
		}

		//mengecek jumlah tiket
		if jmlhtiket >= 5 {
			fmt.Println("Harga: ",hargatiket * jmlhtiket - (hargatiket * jmlhtiket) * 10/100,"\n ")
			total = total + hargatiket * jmlhtiket - (hargatiket * jmlhtiket) * 10/100
		} else if jmlhtiket < 5 && jmlhtiket > 0 {
			fmt.Println("Harga: ",hargatiket * jmlhtiket,"\n ")
			total = total + hargatiket * jmlhtiket
		} else {
			fmt.Println("Harga: ",hargatiket * 1,"\n ")
			total = total + hargatiket * 1
		}

		loop++

		if loop == pesanan {// break
			break
		}
	}
	fmt.Println("Total Harga",total)// total
}