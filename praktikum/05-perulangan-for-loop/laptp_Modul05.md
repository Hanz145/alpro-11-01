# <h1 align="center">Tugas Pendahuluan Modul 05 - Perulangan-For-Loop</h1>
<p align="center">Reihan - 109092600002</p>

### 1. Perulangan Dalam Go
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

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![alt text](https://github.com/Hanz145/alpro-11-01/blob/main/praktikum/05-perulangan-for-loop/tp/Soal-No3/z.output.png?raw=true)


#### Deskripsi
Mengulang program sesuai kondisi lalu menampilkan semua bilangan ganjil dari 1 sampai bilangan yang di masukan atau n lalu menjumlahkan semua bilangan hasil dari perulangan.

### 2. Perulangan Perkalian
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

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/Hanz145/alpro-11-01/blob/main/praktikum/05-perulangan-for-loop/tp/Soal-No4/z.%20output.png?raw=true)


#### Deskripsi
Mengkalikan var `a` hingga mencapai variabel `b` mulai dari `a x 1`

## Kesimpulan
Tujuan pendahuluan kali ini untuk mengetahui bagaimana cara nya menggunakan `for` atau perulangan, apa itu perulangan di dalam go, dan apa fungsi dan manfaat perulangan di go