package main

import "fmt"

func main() {
    var tahun int
	fmt.Scan(&tahun)

	if (tahun % 4 == 0 && tahun % 100 != 0) || (tahun % 400 == 0) {
		fmt.Print(tahun, " masehi merupakan tahun kabisat")
	} else {
		fmt.Print(tahun, " masehi bukanlah tahun kabisat")
	}

}