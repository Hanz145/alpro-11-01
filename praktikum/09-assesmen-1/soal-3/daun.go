package main
import "fmt"

func main(){
	var count, banyak int
	tersempit := 999
	fmt.Scan(&banyak)
	for banyak > 0{
		fmt.Scan(&count)
		if count < tersempit {
			tersempit = count
		}
		banyak --
	}
	fmt.Print(tersempit)
}