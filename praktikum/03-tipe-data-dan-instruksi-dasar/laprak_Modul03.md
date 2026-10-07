# <h1 align="center">Laporan Praktikum Modul 03-Tipe-Data-Dan-Instruksi-Dasar</h1>
<p align="center">Reihan - 109092600002</p>

## Dasar Teori

### A. Tipe Data Dan Instruksi
Tipe data adalah konsep penting dalam pemrograman. Tipe data menentukan ukuran dan jenis nilai variabel.
Go menggunakan pengetikan statis, yang berarti bahwa setelah tipe variabel didefinisikan, variabel tersebut hanya dapat menyimpan data dengan tipe tersebut.
Go memiliki tiga tipe data dasar.

#### 1. Boolean
Boolean adalah tipe data yang hanya bisa menyimpan nilai suatu variabel berupa `true` atau `false`
```go
package main

import "fmt"

func main(){
    //menyimpan variabel a berupa true
    var a bool = true
    //output
    fmt.Print(a)
}
```

#### 2. Numerik
Adalah Tipe data yang bisa menyimpan nilai berupa bilangan. Bisa bilangan bulat maupun bilangan real
```go
package main

import "fmt"

func main(){
    //menyimpan variabel a berupa bilangan bulat
    var a int = 2
    //menyimpan variabel b berupa bilangan real
    var b float64 = 2.5
    //output
    fmt.Print(a)
    fmt.Print(b)
}
```

#### 2. String
Adalah Tipe data yang bisa menyimpan nilai berupa string. Bisa berupa kalimat, maupun satu karakter
```go
package main

import "fmt"

func main(){
    //menyimpan variabel a berupa bilangan bulat
    var a string = "suki"
    //menyimpan variabel b berupa bilangan real
    var b rune = "suki"
    //output
    fmt.Print(a)
    fmt.Print(b)
}
```

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. kasir.go

```go
package main

import "fmt"

func main() {
	var uang int
	fmt.Print("Uang = ")
    //input
	fmt.Scan(&uang)
    //penyelesaian
	sepuluh := uang / 10000
	sisa := uang % 10000
	lima := sisa / 5000
	sisa = sisa % 5000
	satu := sisa / 1000
	sisa = sisa % 1000
    //output
	fmt.Println("10000 = ", sepuluh)
	fmt.Println("5000 = ", lima)
	fmt.Println("1000 = ", satu)
	fmt.Print("sisa = ", sisa)
}
```
#### Deskripsi
Kode program di atas membantu mencari cacah uang. Mencari berapa banyak uang (10k, 5k, 1k) Rupiah, dan sisa uang dengan cara membagi dan memodul total uang 

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
Kode di atas berfungsi mengkonversi suhu dengan satuan celcius ke ke satuan kelvin menggunakan variabel float

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
Kode di atas berfungsi untuk menukarkan tiga nilai variabel sesuat urutan yang di inginkan. `x, y, z` ke `z, y, x`
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
![Screenshot Output Unguided](https://github.com/Hanz145/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/suhu/output.png?raw=true)


#### Deskripsi
Kode di atas bisa mengkonversi suhu dengan satuan celcius ke reamur dengan cara `4.0 / 5.0 * suhu celcius`

### 2. Membuat sebuah program dalam bahasa Go yang dapat mengonversi jumlah hari ke dalam satuan tahun, bulan, minggu, dan hari

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
![Screenshot Output Unguided](https://github.com/Hanz145/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/hari/output.png?raw=true)

#### Deskripsi
Pada kode di atas kita bisa tau berapa tahun, bulan, dan minggu jika kita memasukan jumlah hari. dengan cara membagi dan memodul jumlah hari yang di masukan dengan banyak nya hari yang ada di dalam satu tahun, bulan, dan minggu

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
Berdasarkan 03-tipe-data-dan-instruksi-dasar, saya bisa mempelajari cara mengkonversi, menggunakan modul, dan memindahkan nilai suatu variabel ke variabel lain nya di dalam Golang

## Referensi
1. *w3schools*. Diakses pada 7 Oktober 2026 melalui https://www.w3schools.com/go/
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
