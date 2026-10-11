package main

import "fmt"

func main (){
	var uang, pajak float64
	fmt.Println("====== Masukan uang (sudah berbentuk juta) ======")
	fmt.Print("uang: " ); fmt.Scan(&uang)

	if uang <= 50 {
		pajak = (uang * 5 / 100)
		fmt.Print("pajak: ",pajak, " juta")

	} else if uang > 50 && uang <= 100 {
		pajak = (50.0 * 5.0 / 100.0) +
		((uang - 50.0) * 10.0 / 100.0)
		fmt.Print("pajak: ",pajak, " juta")

	} else if uang > 100 && uang <= 200 {
		pajak = (50.0 * 5.0 / 100.0) +
		(50.0 * 10.0 /100.0) +
		((uang - 100.0) * 15.0 / 100.0)
		fmt.Print("pajak: ",pajak, " juta")

	} else {
		pajak = (50.0 * 5.0 / 100.0) +
		(50.0 * 10.0 /100.0) +
		(100.0 * 15.0 /100.0) +
		((uang - 200.0) * 20.0 / 100.0)
		fmt.Print("pajak: ",pajak, " juta")
	}
}