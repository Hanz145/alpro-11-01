# <h1 align="center">Tugas Pendahuluan Modul [Nomor Modul] - [Judul Modul/Topik]</h1>
<p align="center">[Nama Praktikan] - [NIM]</p>

### 1. Sisa Kue

```go
package main

import "fmt"

func main() {
	var x, y int
	//memasukan nilai
	fmt.Println("Masukan (jumlah anggota, keluarga dan jumlah kue) : ")
	fmt.Scan(&x)
	fmt.Scan(&y)
	//mencari sisa bagi
	sisahasil := y % x
	//sisa bagi
	fmt.Print("Sisa kue : ")
	fmt.Print(sisahasil)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/renwxyz/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instriksi-dasar/tp/sisa/output.png)


#### Deskripsi
Membuat suatu program yang bisa menentukan hasil bagi dari dua variabel. Pada program kita bisa menemukan hasil bagi dari kue yang di bagi ke angora keluarga. Hasil yang diperoleh adalah tau cara menggunakan persen dan menemukan hasil bagi 

### 2. Boolean

```go
package main

import "fmt"

func main() {
	var status bool
	//memasukan (jika true = true, jika false or else = false)
	fmt.Printf("input : ")
	fmt.Scan(&status)
    //output
	fmt.Printf("output : ")
	fmt.Print(status)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/renwxyz/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instriksi-dasar/tp/konversi/output.png)


#### Deskripsi
Dari kode di atas kita tahu cara menggunakan variabel `bool` dengan baik dan benar.

### 2. Boolean

```go
package main

import "fmt"

func main() {
	var mil float64
	var km float64
	//input
	fmt.Printf("Mil : ")
	fmt.Scan(&mil)
	//proses
	km = mil * 1.6
	fmt.Printf("km : ", km)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/renwxyz/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instriksi-dasar/tp/konversi/output.png)

## Kesimpulan
[Tuliskan kesimpulan yang menjawab tujuan praktikum berdasarkan hasil yang diperoleh.]