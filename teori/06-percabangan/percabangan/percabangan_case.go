package main

import "fmt"

func main(){
	x := 11
	y := 10
	fmt.Print(x == 10 || y == 10)
	// output true karena itu di baca atau
	// || -> False || True -> True
}