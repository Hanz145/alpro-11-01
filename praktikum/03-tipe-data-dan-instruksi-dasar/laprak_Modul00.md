# <h1 align="center">Laporan Praktikum Modul [Nomor Modul] - [Judul Modul/Topik]</h1>
<p align="center">[Nama Praktikan] - [NIM]</p>

## Dasar Teori

### A. Tipe Data Dan Instruksi
[Tuliskan penjelasan teori terkait topik ini. Sertakan kutipan/rujukan jika perlu, contoh: Menurut [Nama Penulis] ([Tahun]), ...]

### B. [Judul Topik Dasar Teori 2, misal: Package dan Struktur Program di Go]

#### 1. [Judul Sub-topik 1, misal: Pengertian Package main dan func main()]
[Penjelasan sub-topik 1]

#### 2. [Judul Sub-topik 2, misal: Tipe Data dan Deklarasi Variabel di Go]
[Penjelasan sub-topik 2]

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. kasir.go

```go
package main

import "fmt"

func main() {
	var uang int
	fmt.Print("Uang = ")
	fmt.Scan(&uang)

	sepuluh := uang / 10000
	sisa := uang % 10000
	lima := sisa / 5000
	sisa = sisa % 5000
	satu := sisa / 1000
	sisa = sisa % 1000
	fmt.Println("10000 = ", sepuluh)
	fmt.Println("5000 = ", lima)
	fmt.Println("1000 = ", satu)
	fmt.Print("sisa = ", sisa)
}
```
#### Deskripsi
[Tuliskan ringkasan proses praktikum: apa yang dikerjakan, bagian guided dan unguided yang diimplementasikan, serta hasil yang diperoleh.]

### 2. konversi.go

```go
package main

import "fmt"

func main() {
	var cel float64
	fmt.Print("Masukan Suhu: ")
	fmt.Scan(&cel)

	fmt.Print(cel + 273)
}

```
#### Deskripsi
[Tuliskan ringkasan proses praktikum: apa yang dikerjakan, bagian guided dan unguided yang diimplementasikan, serta hasil yang diperoleh.]

### 3. tukar.go

```go
package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	temp := x
	x = z
	z = y
	y = temp
	fmt.Println(x, y, z)
}
```
#### Deskripsi
[Tuliskan ringkasan proses praktikum: apa yang dikerjakan, bagian guided dan unguided yang diimplementasikan, serta hasil yang diperoleh.]

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. Mengonversi suhu dari derajat Celsius menjadi derajat Reamur

```go
package main

import "fmt"

func main() {
	var celcius float64
	fmt.Println("======= Suhu Celcius ======")
	fmt.Print("Celcius : ")
	fmt.Scan(&celcius)
	reamur := 4.0 / 5.0 * celcius
	fmt.Println("======= Suhu Reamur ======")
	fmt.Print("Reamur : ", reamur)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/renwxyz/alpro-11-01/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/cacahuang/output.png)


#### Deskripsi
[Tuliskan ringkasan proses praktikum: apa yang dikerjakan, bagian guided dan unguided yang diimplementasikan, serta hasil yang diperoleh.]

### 2. Membuat sebuah program dalam bahasa Go yang dapat mengonversi jumlah hari ke dalam satuan tahun, bulan, minggu, dan hari,

```go
package main

import "fmt"

func main() {
	var tahun, bulan, minggu, hari, sisa int32
	fmt.Println("====== Banyak Hari ======")
	fmt.Print("Hari : ")
	fmt.Scan(&hari)
	tahun = hari / 365
	sisa = hari % 365
	bulan = sisa / 30
	sisa = sisa % 30
	minggu = sisa / 4
	sisa = sisa % 4
	fmt.Println("====== Konverenis ke Tahun, Bulan, Minggu ======")
	fmt.Println("Tahun : ", tahun)
	fmt.Println("Bulan : ", bulan)
	fmt.Println("Minggu : ", minggu)
	fmt.Println("sisa : ", sisa)
}

```

##### Output
![Screenshot Output Unguided](unguided/[nama_soal]/output.png)

#### Deskripsi
[Tuliskan ringkasan proses praktikum: apa yang dikerjakan, bagian guided dan unguided yang diimplementasikan, serta hasil yang diperoleh.]

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
[Tuliskan kesimpulan yang menjawab tujuan praktikum berdasarkan hasil yang diperoleh.]

## Referensi
1. [Nama Penulis]. ([Tahun]). *[Judul Buku/Sumber]*. [Kota]: [Penerbit]. Diakses pada [tanggal akses] melalui [tautan/DOI]
2. [Nama Penulis]. ([Tahun]). *[Judul Buku/Sumber]*. [Kota]: [Penerbit]. Diakses pada [tanggal akses] melalui [tautan/DOI]
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
