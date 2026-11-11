package main
import "fmt"
func main(){
	var m, k, t int = 0, 0, 0
	var check string
	for {
		fmt.Scan(&check)
		if check == "T"{
			t ++
		} else if check == "M" {
			m ++
		} else if check == "K" {
			k ++
		} else {
			break
		}
	}
	fmt.Println("Tipe M = ", m)
	fmt.Println("Tipe K = ", k)
	fmt.Println("Tipe T = ", t)
}