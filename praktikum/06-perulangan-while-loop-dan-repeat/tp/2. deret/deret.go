package main

import "fmt"

func main(){
	var n int
	hasil := 0
	loop := 0
	fmt.Scan(&n)
	for loop <= n{
		hasil = hasil + loop
		loop ++
	}
	fmt.Print(hasil)
}