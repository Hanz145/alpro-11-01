package main
import "fmt"

func hitung(x,y int /*untukmenyatakan var int*/) int /*untuk mengeluarkan int*/ {
	return 2+x+y-5
}

func main(){
	fmt.Println(hitung(2,5)) // mendeklarasikan
	fmt.Println(hitung(2,10)) // mendeklarasikan
}