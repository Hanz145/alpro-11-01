package main

import "fmt"

func main (){
	var clo string
	var harga int
	fmt.Scan(&harga)
	fmt.Scan(&clo)

	if clo == "true"{
		fmt.Print(harga - (harga * 35/100))
	} else {
		fmt.Print(harga)
	}
}
