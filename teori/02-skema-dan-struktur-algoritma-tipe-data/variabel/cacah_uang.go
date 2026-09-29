package main
import "fmt"
func main(){
	var uang, sisa int
	fmt.Scan(&uang)
	fmt.Println(uang,"uang / 10000")
	fmt.Println("lembar")
	sisa = uang % 10000
}

