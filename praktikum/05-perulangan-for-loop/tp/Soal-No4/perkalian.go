package main

import "fmt"

func main(){
	var a, b, hasil, c int
	fmt.Scan(&a)
	fmt.Scan(&b)

	c = 0

	for loop := 0; loop < b; loop++{
		c++
		hasil = a * c
		fmt.Println(a," x ",c," = ",hasil)
	}
}