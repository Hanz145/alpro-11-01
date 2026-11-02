# <h1 align="center">Tugas Pendahuluan Modul 06 - Perulangan-While-Loop-dan-Repeat</h1>
<p align="center">Reihan - 109092600002</p>

### 1. Lulus.go
```go
package main


import "fmt"

func main(){
	var nilai int // var nilai

	for {
		fmt.Print("Masukan Nilai: ")
		fmt.Scan(&nilai)
		if nilai >= 60 && nilai < 101{
			fmt.Println("Anda lulus ujian")
		} else if nilai < 60 && nilai > -1{
			fmt.Println("Anda tidak lulus ujian")
		} else if nilai == -1 {
			break
		} else {
			fmt.Println("tidak ada nilai, mohon masukan angka yang valid (0 - 100)")
		}
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![alt text](https://github.com/Hanz145/alpro-11-01/blob/main/praktikum/06-perulangan-while-loop-dan-repeat/tp/1.%20lulus/z.output.png?raw=true)


#### Deskripsi
Menentukan apakah mahasiswa lulus atau tidak dan menghitung berapa banyak mahasiswa yang lulus sesuai banyak nya mahasiswa yang di masukan.

### 2. Deret.go
```go
package main

import "fmt"

func main(){
	var n int
	hasil := 0
	loop := 0
	fmt.Scan(&n)
	for loop <= n{
		hasil = hasil + loop
		loop ++
	}
	fmt.Print(hasil)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/Hanz145/alpro-11-01/blob/main/praktikum/06-perulangan-while-loop-dan-repeat/tp/2.%20deret/z.output.png?raw=true)


#### Deskripsi
Menambahkan bilangan dari `1` sampai `n` (bilangan yang di input)

### 3. Dua Terendah
```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var namaTerendah, namaKeduaTerendah string
	var nilaiTerendah, nilaiKeduaTerendah int

	nilaiTerendah = 100
	nilaiKeduaTerendah = 100

	var i int = 0
	for i < n {
		var nama string
		var nilai int
		fmt.Scan(&nama, &nilai)

		if nilai < nilaiTerendah {
			nilaiKeduaTerendah = nilaiTerendah
			namaKeduaTerendah = namaTerendah
			nilaiTerendah = nilai
			namaTerendah = nama
		} else if nilai < nilaiKeduaTerendah {
			nilaiKeduaTerendah = nilai
			namaKeduaTerendah = nama
		}

		i++
	}
	fmt.Println(namaTerendah)
	fmt.Println(namaKeduaTerendah)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/Hanz145/alpro-11-01/blob/main/praktikum/06-perulangan-while-loop-dan-repeat/tp/3.%20dua-terendah/z.output.png?raw=true)


#### Deskripsi
Program di atas membantu mencari dua mahasiswa dengan nilai terendah dari banyak nya mahasiswa yang di input.

### 4. biaya.go
```go
package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    var biaya float64 = 10000000.0
    var i int = 1
    for i < n {
        biaya = biaya * 1.06
        i++
    }
    fmt.Printf("%.0f\n", biaya)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/Hanz145/alpro-11-01/blob/main/praktikum/06-perulangan-while-loop-dan-repeat/tp/4.%20biayaya/z.output.png?raw=true)


#### Deskripsi
Kode di atas berfungsi untuk mencari berapa banyak uang yang harus di bayar jika setiap tahun bertambah 6%.

## Kesimpulan
Tujuan tugas pendahuluan minggu ini yaitu supaya mahasiswa yahu cara menggunakan perulangan while loop dan repeat until.