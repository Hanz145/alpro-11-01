# <h1 align="center">Laporan Praktikum Modul 02 - Variabel Dan Tipe Data Di Dalam Golang</h1>
<p align="center">[Reihan] - [109092600002]</p>

## Dasar Teori

### A. Bahasa Pemrograan Go (Golang)
Golang adalah adalaah bahasa pemrograman open-source yang di kembangkan oleh Google pada tahun 2007 oleh Robert Griesemer , Rob Pike , dan Ken Thompson. Kemudian Golaang di rilis pada November 2009
### B. Variabel, Tipe Data, Dan Operator Dalam Pemrograman Go

#### 1. Variabbel
Di dalam pemrograman Go kita bisa menyimpan data dan biasanya sering di ubah saat pemrograman berjalan. Bisa berupa bilangan bulat, bilangan real, karakter, kondisi, dan juga berupa kata atau kalimat.

#### 2. Tipe Data
Tipe data berfungsi untuk menentukan jenis data apa yang ingin disimpan di dalam suatu pemrograman. **Int** untuk menyimpan bilangan bulat, **float** untuk menyimpan bilangan real, **string** untuk menyimpan suatu kalimat atau kata, **rune** untuk menyimpan suatu karakter, **bool** untuk menyimpan kondisi (benar/salah)
```go
func main(){
	var a int = 67 //bilangan bulat
	var b float64 = 6.7 //bilangan real
	var c string = "Six Seven" //menyimpan kata
	var d rune = "@" //menyimpan suatu karakter
	var e bool = true
}
```

#### 3. Operator
Operator berfungsi untuk melakukan operasi terhadap nilai/variabel di dalam suatu pemrograman.
```go
func main(){
	fmt.Println(a + b) //penjumlahan
	fmt.Println(a - b) //pengurangan
	fmt.Println(a * b) //perkalian
	fmt.Println(a / b) //pembagian
	fmt.Println(a + b) //penjumlahan
	fmt.Println(a % b) //sisa bagi
}
```

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. lingkaran.go

```go
package main
import "fmt"

func main(){
    var pi float32 = 3.14
	var r float32
	fmt.Println("========= Lingkaran =========")
	fmt.Print("Masukan jari jari lingkaran: ")
	fmt.Scan(&r)
	luas := pi * r * r
	fmt.Println("========= Luas Lingkaran =========")
	fmt.Println("Luas = ", pi, " * ", r, " * ", r, " = ", luas)
}
```
#### Deskripsi
Pada kode pemrograman di atas, kita bisa mencari luas lingkaran dengan hanya memasukan jari-jari sebuah lingkaran. Cara kerja code ini yaitu memasukan jari-jari lalu code ini akan mengalikan jari-jari nya dua kali dan mengalikan nya lagi dengan variabel pi (3.14) sehingga mendapatkan hasil luas lingkaran

### 2. skor.go

```go
package main

import "fmt"

func main() {
	//varibel
	var NamaSiswa string
	var NilaiBahasaInggris, NilaiMatematika int
	//input
	fmt.Print("Nama: ")
	fmt.Scan(&NamaSiswa)

	fmt.Print("Nilai Bahasa Inggris: ")
	fmt.Scan(&NilaiBahasaInggris)

	fmt.Print("Nilai Matematika: ")
	fmt.Scan(&NilaiMatematika)
	//rata rata
	jumlah := NilaiBahasaInggris + NilaiMatematika
	ratarata := jumlah / 2
	//keluaran
	fmt.Print("Rata-Rata: ", ratarata)
}
```
#### Deskripsi
Pada kode pemrograman di atas, kita bisa mencari rata-rata nilai seorang siswa. dengan memasukan nama, nilai bahasa ingris siswa, dan nilai matematika siswa, nanti program akan menjumlahkan nilai matematika dan bahasa ingris lalu membagi nya dengan dua (karena cuman ada dua nilai yang di jumlahkan). Keluaran nya, si siswa akan mendapatkan rata-rata nilai yang akurat

### 3. suhu.go

```go
package main

import "fmt"

func main() {
	var celsius float32
	fmt.Println("======== Suhu =========")
	fmt.Print("Celsius: ")
	fmt.Scan(&celsius)
	reamur := celsius * 4.0 / 5.0
	fahrenheit := celsius*9.0/5.0 + 32.0
	kelvin := celsius + 273.15
	fmt.Println("======== Hasil Suhu =========")
	fmt.Println("reamur = ", reamur)
	fmt.Println("fahrenheit = ", fahrenheit)
	fmt.Println("kelvin = ", kelvin)
}
```

#### Deskripsi
Kode pemrograman di atas berfungsi mengubah suhu dengan satuan selsius menjadi **reamur** (celsius * 4.0 / 5.0), **farenheit** (celsius * 9.0/5.0 + 32.0), dan **kelvin** (celsius + 273.15)

### 4. tukar.go

```go
package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Aku adalah pesulap yang bisa menukarkan nilai kedua variabel. Silahkan masukan angka nya Tuan/nyonya")
	fmt.Print("Variabel A: ")
	fmt.Scan(&a)
	fmt.Print("Variabel B: ")
	fmt.Scan(&b)
	c = a
	a = b
	b = c
	fmt.Println("Ini hasil nya tuan")
	fmt.Println("Variabel A :", a)
	fmt.Println("Variabel B :", b)
}
```
#### Deskripsi
Kode pemrograman di atas berfungsi untuk menukar nilai variabel **A** dan **B** dengan cara membuat variabel **C** dan memasukan nilai  variabel **A** sebagai jembatan. Lalu membuat nilai variabel **A** sama dengan variabel **B**. Lalu membuat nilai variabel **B** saama dengan variabel **C**. Kemudian yang akan keluar (output) adalah **A** = **B** dan **B** = **A**

## Unguided

### 1. Cacah Uang

```go
package main

import "fmt"

func main() {
	var jumlahuang int32
	fmt.Println("=========Jumlah Uang=========")
	fmt.Print("Masukan Jumlah Uang: ")
	fmt.Scan(&jumlahuang)
	uang10k := jumlahuang / 10000
	sisauang := jumlahuang % 10000
	uang5k := sisauang / 5000
	sisauang = sisauang % 5000
	uang1k := sisauang / 1000
	sisauang = sisauang % 1000
	fmt.Println("=========Pecahan Uang=========")
	fmt.Println("Uang 10k: ", uang10k)
	fmt.Println("Uang 5k: ", uang5k)
	fmt.Println("Uang 1k: ", uang1k)
	fmt.Println("Uang sisa: ", sisauang)
	fmt.Println("Jumlah uang ", jumlahuang)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/Hanz145/alpro-11-01/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/Cacah%20uang/output.png?raw=true)


#### Deskripsi
Pada kode pemrograman di atas kita bisa mecari cacah uang dalam pecahan sepuluh ribu, lima ribu, dan seribu. Cara kerja nya yaitu dengan memasukan nilai uang yang kita inginkan. Lalu kode pemrograman akan melakukan perhitungan dengan cara nilai uang di bagi sepuluh ribu dan di modul dengan sepuluh ribu. lalu hasil sisa bagi tersebut di kalikan dan di modulkan dengan lima ribu. lakukan hal yang sama di seribu. Output yang keluar adalah hasil bagi dari jumlah uang di bagi sepuluh ribu, sisa hasil bagi sepuluh ribu di bagi lima ribu, lalu hasil bagi lima ribu di bagi seribu. Sebagai tambahan aku menambahkan sisa uang

### 2. Kalkulator

```go
package main

import "fmt"

func main() {
	var a, b int
	for i := 1; i < 2; i = i + 0 {
		fmt.Println("========= Masukan Dua Angka =========")
		fmt.Print("Bilangan A: ")
		fmt.Scan(&a)
		fmt.Print("Bilangan B: ")
		fmt.Scan(&b)
		if b == 0 {
			fmt.Println("maaf, B tidak bisa berupa 0")
		} else {
			break
		}
	}
	tambah := a + b
	kurang := a - b
	kali := a * b
	bagi := a / b
	persen := a % b
	fmt.Println("========= Hasil Bilangan =========")
	fmt.Println("a + b = ", tambah)
	fmt.Println("a - b = ", kurang)
	fmt.Println("a * b = ", kali)
	fmt.Println("a / b = ", bagi)
	fmt.Println("a % b = ", persen)
}
```

##### Output
![Screenshot Output Unguided](https://github.com/Hanz145/alpro-11-01/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/Kalkulator/Output.png?raw=true)

#### Deskripsi
Pada kode pemrograman di atas kita bisa mengetahui hasil dari penjumlahan, pengurangan, perkalian, pembagian, dan hasil pembagian dengan hanya memasukan dua bilangan bulat (a dan b). Namun jika memasukan angka 0 di variabel **B**, akan terjadi error. maka dari itu jika **B** = 0 maka program akan mengulang

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
Berdasarkan praktikum 02-bahasa-pemrograman-go

## Referensi
1. Fika Ridaul Maulayya. (2026). *Belajar Golang Dasar #4: Variable*. Diakses pada 27 September 2026 melalui https://santrikoding.com/belajar-golang-dasar-4-variable

2.  MODUL 2 BAHASA PEMROGRAMAN GO 
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
