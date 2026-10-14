## 2. Bahasa Go tidak memiliki perulangan while dan do-while. Jelaskan perbedaan antara for dengan tiga komponen dan for yang hanya menggunakan kondisi! Berikan contoh kode masing-masing. 
### Jawaban:
For dengan tiga komponen memiliki inisiasi, kondisi, dan post statment setelah if. Sedangkan for dengan kondisi hanya memasukan kondisi setelah if.
####  - For dengan Tiga Komponen
```go
package main

import "fmt"

func main(){
    var n int
    for loop := 0; loop < 5; loop++{
        n = n + 1
        fmt.Print(n)
    }
}
```
####  - For dengan Hanya Kondisi
```go
package main

import "fmt"

func main(){
    var n int
    var loop int
    for loop < 5{
        n = n + 1
		loop = loop + 1
        fmt.Print(n)
    }
}
```