package main

import (
	"fmt"
	"math"
)

func main(){
	var berat float64
	var zona int
	var layanan string

	fmt.Scan(&berat, &zona, &layanan)

	if berat > 30{
		fmt.Println("Berat melebihi batas max (30kg)")
		return
	}

	var tarifperkg int

	switch zona{
		case 1:
			tarifperkg = 10000
		case 2:
			tarifperkg = 15000
		case 3:
			tarifperkg = 25000
		default :
			fmt.Println("zona tidak valid")
			return
	}

	beratbulat := int(math.Ceil(berat))
	total := beratbulat * (tarifperkg)

	if layanan == "Express"{
		total = total + total / 2
	} else if layanan != "Reguler"{
		fmt.Println("Layanan tidak valid")
		return
	}

	fmt.Println(beratbulat)
	fmt.Println(total)
}