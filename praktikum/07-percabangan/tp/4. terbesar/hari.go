package main
import "fmt"
func main() {
	var a, b, c, max int
	fmt.Scan(&a, &b, &c)

	if a >= b && a >= c {
		max = a
	} else if b >= a && b >= c {
		max = b
	} else {
		max = c
	}

	fmt.Println(max)
}