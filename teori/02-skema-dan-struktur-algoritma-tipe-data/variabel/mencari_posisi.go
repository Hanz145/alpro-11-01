package main

import "fmt"

func main() {
	var posisi, posisi1, kecepatan, waktu int
	fmt.Scan(&posisi, &kecepatan, &waktu)
	posisi1 = posisi + (kecepatan * waktu)
	fmt.Println(posisi1)
}
