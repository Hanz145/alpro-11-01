## 4. Buatlah program dalam bahasa Go menggunakan perulangan for untuk menampilkan tabel perkalian dari sebuah bilangan a, mulai dari a × 1 hingga a × b. 
### Jawaban:
```go
package main

import "fmt"

func main(){
	var a, b, hasil, c int
	fmt.Scan(&a)
	fmt.Scan(&b)

	c = 0

	for loop := 0; loop < b; loop++{
		c++
		hasil = a * c
		fmt.Println(a," x ",c," = ",hasil)
	}
}
```