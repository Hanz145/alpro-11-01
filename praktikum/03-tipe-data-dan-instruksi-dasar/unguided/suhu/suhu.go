package main

import "fmt"

func main() {
	var celcius float64
	fmt.Println("======= Suhu Celcius ======")
	fmt.Print("Celcius : ")
	fmt.Scan(&celcius)
	reamur := 4.0 / 5.0 * celcius
	fmt.Println("======= Suhu Reamur ======")
	fmt.Print("Reamur : ", reamur)
}
