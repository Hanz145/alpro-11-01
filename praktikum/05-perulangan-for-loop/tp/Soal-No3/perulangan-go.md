## 3.Buatlah program dalam bahasa Go menggunakan perulangan for untuk menampilkan semua bilangan ganjil dari 1 hingga n, kemudian menampilkan jumlah seluruh bilangan ganjil tersebut. 
### Jawaban:
```go
package main

import "fmt"

func main(){
	var n1 int = 1
	var bil, hasil int
	fmt.Scan(&bil)
	for n1 <= bil{
		fmt.Print(" ",n1)
		hasil = hasil + n1
		n1 = n1 + 2
	}
	fmt.Print("\n",hasil)
}
```