package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("Nol")
	} else {
		var jenis, status string

		if n%2 == 0 {
			jenis = "Genap"
		} else {
			jenis = "Ganjil"
		}

		if n > 0 {
			status = "Positif"
		} else {
			status = "Negatif"
		}

		fmt.Println(jenis + ", " + status)
	}
}