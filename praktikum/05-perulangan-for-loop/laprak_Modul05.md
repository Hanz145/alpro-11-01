# <h1 align="center">Laporan Praktikum Modul 05-Perulangan-For-Loop</h1>
<p align="center">Reihan - 109092600002</p>

## Dasar Teori

### A. Loop Di Dalam Go
Go hanya memiliki `for` loop yang berarti Go tidak memiliki `while` loop seperti bahasa pemrograman C++. Tidak seperti bahasa lain seperti C, Java, atau JavaScript, tidak ada tanda kurung yang mengapit ketiga komponen pernyataan tersebut fordan tanda kurung kurawal `{ }` selalu diperlukan. Loop Dalam Bahasa Go memiliki tiga komponen penting.
#### 1. **Inisialisasi**
adalah komponen yang di eksekusi sekali sebelum perulangan. Berfungsi untuk menghitung dan mengatur perulangan.
#### 2. **Kondisi**
adalah ekspresi boolean yang dievaluasi sebelum setiap perulangan dimulai. Berfungsi untuk menentukan apakah perulangan harus di lanjutkan atau di hentikan.
#### 3. **Post statment**
adalah pernyataan yang dieksekusi di akhir setiap iterasi. Berfungsi untuk mengubah suatu variabel dan menampilkan output.

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. tabungan.go

```go
package main

import "fmt"

func main(){
	var a, b, n int
	fmt.Scan(&a, &b, &n)
	total := 0
	tabunganhariini := a

	for i := 1; i <= n; i++{
		total = total + tabunganhariini
		tabunganhariini = tabunganhariini + b
	}
	fmt.Print(total)
}
```

#### Deskripsi
Kode di atas menampilkan jumlah tabungan yang terkumpul sesuai hari yang di tentukan. Pertama kita di suruh menginput `a` adalah tabungan hari pertama, `b` adalah yang akan di tabung setiap hari, dan `c` adalah berapa hari kita akan menabung. Output akan menampilkan jumlah tabungan di hari yang telah di input.

### 2. terbesar.go
```go
package main

import "fmt"

func main() {

	var hari, tiket string // var hari sama jenis tiket
	var pesanan, total, jmlhtiket int // var banyak nya pesanan dan total harga pesanan
	fmt.Println("Silahkan Pesan tiket bioskop yang kamu mau")
	fmt.Println("==========================")
	fmt.Println("Reguler = 40000")
	fmt.Println("3D = 60000")
	fmt.Println("IMAX = 80000")
	fmt.Println("Jumat - Minggu = +10000")
	fmt.Println("==========================")
	fmt.Println("diskon 10% jika membeli 5 tiket dalam sehari")
	fmt.Println("jika ada kesalahan dalam memasukan hari, maka hari = Senin")
	fmt.Println("jika ada kesalahan dalam nama tiket, maka tiket = Reguler")
	fmt.Println("jika ada kesalahan dalam memasukan jumlah tiket, maka tiket = 1")
	fmt.Print("Banyak Pesanan: ")
	fmt.Scan(&pesanan)// berapa banyak tiket yang akan di pesan
	loop := 0 //var pengatur loop
	fmt.Print("\n ")


	for { // loop utama
		fmt.Println("====== Tiket ======")
		fmt.Print("Jenis Tiket: ")
		fmt.Scan(&tiket)
		fmt.Print("Hari (diawali kapital): ")
		fmt.Scan(&hari)
		fmt.Print("Jumlah Tiket: ")
		fmt.Scan(&jmlhtiket)
		hargatiket := 0 //menampilkan harga sesuai jenis tiket

		//mengecek harga tiket
		if tiket == "IMAX"{
			hargatiket = 80000
		} else if tiket == "3D" {
			hargatiket = 60000
		} else {
			hargatiket = 40000
		}

		//mengecek hari
		if hari == "Jumat" || hari == "Sabtu" || hari == "Minggu" {
			hargatiket += 10000
		} else if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
			hargatiket += 0
		} else {
			hargatiket += 0
		}

		//mengecek jumlah tiket
		if jmlhtiket >= 5 {
			fmt.Println("Harga: ",hargatiket * jmlhtiket - (hargatiket * jmlhtiket) * 10/100,"\n ")
			total = total + hargatiket * jmlhtiket - (hargatiket * jmlhtiket) * 10/100
		} else if jmlhtiket < 5 && jmlhtiket > 0 {
			fmt.Println("Harga: ",hargatiket * jmlhtiket,"\n ")
			total = total + hargatiket * jmlhtiket
		} else {
			fmt.Println("Harga: ",hargatiket * 1,"\n ")
			total = total + hargatiket * 1
		}

		loop++

		if loop == pesanan {// break
			break
		}
	}
	fmt.Println("Total Harga",total)// total
}
```

#### Deskripsi
Kode di atas berfungsi untuk menampilkan harga tiket sesuai kondisi dan juga menampilkan total semua harga tiket yang telah di beli. Di dalam kode ini juga kita bisa mengatur seberapa banyak tiket yang ingin kita pesan.
##### Catatan: Aku sedikit mengubah tampilan/output namun tidak mengubah fungsi nya sama sekali. Ini aku lakukan untuk memperjelas fungsi kode dan output nya.

## Unguided
### 1. Buatlah program dalam bahasa Go untuk menampilkan seluruh faktor dari sebuah bilangan bulat positif n, kemudian menentukan apakah n merupakan bilangan prima atau bukan
```go
package main

import "fmt"

func main (){
	var n, hasil int // var n dan hasil
	var count int // var pendeteksi

	fmt.Println("====== Bilaangan Prima ======")
	i := 1
	fmt.Print("Bilangan: ")
	fmt.Scan(&n) //input
	limit := n

	fmt.Print("Faktor: ")
	for { // perulangan 1
		hasil = n % i
		if hasil == 0 { // jika n == 0
			fmt.Print(i," ")
			count += 1
		}
		i += 1
		if i > limit { // break
			break
		}
	}

	fmt.Print("\nPrima: ")
	if count < 3 {
		fmt.Print("true") // jika prima maka hanya mengitung dua angka
	} else {
		fmt.Print("False") // jika tidak dia akan menghitung lebih dari dua
	}
}
```
##### Output
![Screenshot Output Unguided](https://github.com/Hanz145/alpro-11-01/blob/main/praktikum/05-perulangan-for-loop/unguided/1-prima/z.output.png?raw=true)

#### Deskripsi
Kode di atas berfungsi untuk menampilkan berapa banyak bilangan yang bisa di bagi dengan bilangan yang telah kita masukan. jika bilangan yang di masukan hanya bisa di bagi oleh dua bilangan berarti bilangan prima. Jika tidak maka bukan bilangan prima.

### 2. Buatlah program dalam bahasa Go untuk menentukan apakah setiap mahasiswa baru Tel-U dalam satu kelompok bisa mendapat sertifikat ECCT atau tidak.

```go
package main

import "fmt"

func main(){

	var siswa int // var mengatur banyak siwswa
	fmt.Println("====== ECCT Siswa ======")
	fmt.Print("Jumlah Siswa: ")
	fmt.Scan(&siswa) // input
	loop := 0 // var loop
	var  P, V, GA, FS, CI, TC, ratarata float32 //untuk nilai
	count := 1
	countlulus := 0 // menghitung siswa lulus

	for { // loop
		fmt.Println("====== Nilai Siswa ======")
		fmt.Println("Siswa: ", count)
		fmt.Print("Nilai P: ")
		fmt.Scan(&P)

		fmt.Print("Nilai V: ")
		fmt.Scan(&V)

		fmt.Print("Nilai GA: ")
		fmt.Scan(&GA)

		fmt.Print("Nilai FS: ")
		fmt.Scan(&FS)

		fmt.Print("Nilai CI: ")
		fmt.Scan(&CI)

		fmt.Print("Nilai TC: ")
		fmt.Scan(&TC)

		ratarata = (P + V + GA + FS + CI + TC)/6

		fmt.Println("Nilai Minimal Rata-Rata: 3,33")
		fmt.Println("Rata-Rata: ", ratarata)

		if ratarata >= 3.33 {
			fmt.Println("Status: Ya") // jika di atas rata rata
			countlulus += 1
		} else {
			fmt.Println("Lulus: Tidak") // jika tidak
		}
		
		loop += 1
		if loop == siswa{
			break
		}
	}
	fmt.Print("Sisiwa Yang Lulus: ",countlulus)
}
```

##### Output
![Screenshot Output Unguided](https://github.com/Hanz145/alpro-11-01/blob/main/praktikum/05-perulangan-for-loop/unguided/2-ect/z.output.png?raw=true)

#### Deskripsi
Kode di atas berfungsi untuk menentukan apakah seorang mahasiswa siswa atau beberapa mahasiswa mendapatkan sertifikan ECCT atau tidak dengan cara mencari nilai rata-rata dari hasil test nya. Apabila nilai rata-rata di atas atau setidaknya 3.33, maka mahasiswa di nyatakan lulus test. Kode juga menampilkan berapa banyak mahasiswa yang lulus di dalam tes tersebut. 

## Referensi
1. *A Tour Of Go*. Diakses pada 18 Oktober 2026 melalui https://go.dev/tour/flowcontrol/1
2. *GitHub*. Diakses pada 18 Oktober 2026 melalui https://github.com/Hanz145/alpro-11-01/blob/main/praktikum/05-perulangan-for-loop/tp/Soal-No1/fungsi-utama-struktur-for.md
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
