package main

import "fmt"

func main(){
	var n1 int = 1
	var bil, hasil int
	fmt.Scan(&bil)
	for n1 <= bil{
		fmt.Print(" ",n1)
		hasil = hasil + n1
		n1 = n1 + 2
	}
	fmt.Print("\n",hasil)
}