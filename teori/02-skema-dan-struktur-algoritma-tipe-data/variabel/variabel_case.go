package main

import "fmt"

func main() {
	var name string
	name = "Abdul Suki Liar"
	fmt.Println("nama: ", name)

	var lastname string = "Liar"
	fmt.Println("Nama Belakang: ", lastname)

	middlename := "Suki"
	fmt.Println("Nama Tengah: ", middlename)
	var (
		fullname  = "Reihan aja"
		firstname = "Reihan"
	)
	fmt.Println("nama lengkap:", fullname)
	fmt.Println("Nama depan: ", firstname)
}
