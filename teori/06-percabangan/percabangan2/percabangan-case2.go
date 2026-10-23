package main

import "fmt"

func main(){
	lamp_stat := "Kuning" // var

	if lamp_stat == "Merah"{ // menentukan output sesuai kondisi var
		fmt.Print("Berhenti")
	} else if lamp_stat == "Kuning" {
		fmt.Print("Hati-Hati")
	} else if lamp_stat == "Jalan" {
		fmt.Print("Jalan")
	} else {
		fmt.Print("Perintah tidak terdaftar")
	}
}
