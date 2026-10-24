package main

import "fmt"

func main(){
	var opsi int
	var done bool = false
	pw := 1234
	uang := 10000000
	try := 3
	

	fmt.Println("====== Masukan Pasword ======")

	//loop utama
	for {
		if done {
			break
		}
		ins_pw := 0000
		fmt.Print("Pasaword: ")
		fmt.Scan(&ins_pw)

		//jika pw salah
		if ins_pw != pw {
			fmt.Println("Pasword yang anda masukan salah")
			try = try - 1
			
			//jika jumlah percobaan mencapai max
			if try == 0{
				fmt.Println("Maaf ATM Anda Terkunci")
				break
			}

			//jika atm berhasil
		} else {
			fmt.Println("Login Berhasil")

			//loop jika berhasil
			for {
				opsi = 0
				tarik := 0
				fmt.Println("\n====== Bank ATM ======")
				fmt.Println("1. Lihat Saldo")
				fmt.Println("2. Tarik Uang")
				fmt.Println("3. Keluar")
				fmt.Print("Opsi: ")
				fmt.Scan(&opsi)

				//opsi 1
				if opsi == 1 {
					fmt.Println("Saldo: Rp", uang)

				//opsi 2
				} else if opsi == 2 {
					fmt.Print("Jumlah Uang: ")
					fmt.Scan(&tarik)
					if tarik > uang {
						fmt.Println("Saldo anda tidak mencukupi")
					} else if tarik <= 0 {
						fmt.Println("Jumlah Tidak Valid")
					} else {
						uang = uang - tarik
						fmt.Println("Saldo Anda: Rp",uang)
					}
				
				//opsi 3
				} else if opsi == 3{
					done = true
					break

				//opsi lain nya
				} else {
					fmt.Println("Opsi tidak valid")
				}
			}
		}
	}
}