package main

import "fmt"

func main() {
	var mil float64
	//input
	fmt.Printf("Mil : ")
	fmt.Scan(&mil)
	//proses
	km := mil * 1.6
	fmt.Printf("km : ", km)
}
