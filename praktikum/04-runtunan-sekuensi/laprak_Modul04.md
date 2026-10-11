# <h1 align="center">Laporan Praktikum Modul 04-Runtunan-Sekuensi</h1>
<p align="center">Reihan - 109092600002</p>

## Dasar Teori

### A. Switch
Sebuah `switch` pernyataan adalah cara yang lebih singkat untuk menulis serangkaian if - elsepernyataan. Pernyataan ini menjalankan kasus pertama yang nilainya sama dengan ekspresi kondisi. 

### B. If,  Else, dan Else If
pernyataan `if`, `else`, dan `else if` mirip dengan perulangan `for`, pernyataan ini digunakan untuk menentukan perintah mana yang akan di jalankan sesuai kondisi, bisa menjalankan salah satu dari banyak perintah dan sebagian dari banyak perintah


<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. grade.go

```go
package main

import "fmt"
import "bufio"
import "os"

func main() {
	var nama string
	var nilai float32
	fmt.Println("====== Program Penilaian ======")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan nama: ")
	scanner.Scan()
	nama = scanner.Text()
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	fmt.Println("====== Hasil Penilaian ======")
	fmt.Println("Nama:", nama)
	fmt.Println("Nilai:", nilai)

	if nilai >= 90 && nilai <= 100 {
		fmt.Println("Nilai A")
	}
	if nilai >= 80 && nilai < 90 {
		fmt.Println("Nilai B")
	}
	if nilai >= 70 && nilai < 80 {
		fmt.Println("Nilai C")
	}
	if nilai >= 60 && nilai < 70 {
		fmt.Println("Nilai D")
	}
	if nilai < 60 {
		fmt.Println("Nilai F")
	}
}
```

#### Deskripsi
Kode di atas menentukan nilai `A, B, C, D, F` apa yang di dapat dari seorang siswa sesuai kondisi yang ada di dalam program menggunakan `if`.

### 2. klasifikasi.go
```go
package main

import (
	"fmt" // Hanya memerlukan fmt untuk keperluan input dan output
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var usia int
	var gaji int
	var keterangan string

	// Membaca dua masukan berupa angka dari pengguna (usia dan gaji)
	// fmt.Scan otomatis memisahkan input berdasarkan spasi atau baris baru (enter)
	fmt.Scan(&usia)
	fmt.Scan(&gaji)

	// Menggunakan switch tanpa ekspresi.
	// Cara kerjanya sama persis seperti deretan if - else if.
	// Program akan mengecek dari atas ke bawah, dan menjalankan case pertama yang bernilai benar (true).
	switch {
	case usia < 18:
		keterangan = "Masih sekolah"
	case usia >= 18 && usia <= 25 && gaji >= 50:
		keterangan = "Muda sukses"
	case usia >= 18 && usia <= 25 && gaji < 50:
		keterangan = "Masih belajar hidup"
	case usia >= 26 && usia <= 40 && gaji >= 100:
		keterangan = "Pekerja mapan"
	case usia >= 26 && usia <= 40 && gaji < 100:
		keterangan = "Perlu perbaikan karier"
	case usia > 40 && gaji >= 150:
		keterangan = "Profesional berpengalaman"
	case usia > 40 && gaji < 150:
		keterangan = "Perlu evaluasi finansial"
	}

	// Menampilkan hasil klasifikasi ke layar
	fmt.Println(keterangan)
}
```
#### Deskripsi
Kode di atas berdungsi untuk mengklasifikasikan seseorang berdasarkan gaji dan umur sesuai kondisi yang memenuhi di dalam pemrograman menggunakan pernyataann `case`

### 3. penilaian.go
```go
package main

import (
	"bufio" // Digunakan untuk membaca masukan string yang memiliki spasi (seperti nama lengkap)
	"fmt"
	"os" // Menyediakan akses ke sistem operasi, dalam hal ini os.Stdin yang merepresentasikan keyboard
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var pilihan int
	var nama string
	var nilai float64
	var grade string


	// Menampilkan cetakan Menu ke layar
	fmt.Println("============= Menu =============")
	fmt.Println("1. Sistem penilaian")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih opsi: ")

	// Membaca masukan pilihan (angka).
	// Kita menggunakan Scanln agar saat user menekan 'Enter', karakter enter tersebut
	// ikut diolah/dibersihkan dan tidak melompat (mengganggu) masukan nama di bawahnya.
	fmt.Scanln(&pilihan)

	// Percabangan/Sekuensi bersyarat berdasarkan input pengguna
	if pilihan == 1 {
		fmt.Print("Masukkan nama siswa: ")

		// Membuat alat pembaca (scanner) baru yang mengambil masukan dari keyboard
		scanner := bufio.NewScanner(os.Stdin)

		// Proses membaca masukan dari pengguna sampai tombol Enter ditekan
		scanner.Scan()

		// Mengambil teks (nama) yang baru saja dibaca dan menyimpannya ke variabel
		nama = scanner.Text()

		fmt.Print("Masukkan nilai siswa: ")
		fmt.Scanln(&nilai)

		// Menentukan grade nilai
		if nilai >= 90 && nilai <= 100 {
			grade = "A"
		} else if nilai >= 80 && nilai < 90 {
			grade = "B"
		} else if nilai >= 70 && nilai < 80 {
			grade = "C"
		} else if nilai >= 60 && nilai < 70 {
			grade = "D"
		} else {
			grade = "F"
		}

		// %s adalah format (placeholder) untuk mencetak data bertipe string
		// %s pertama akan digantikan oleh isi variabel 'nama', %s kedua oleh 'grade'
		fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)

	} else if pilihan == 0 {
		// Dijalankan jika pengguna mengetik 0
		fmt.Println("Keluar dari program.")

	} else {
		// Dijalankan jika pengguna mengetik angka selain 1 dan 0
		fmt.Println("Pilihan tidak valid.")
	}
}
```

#### Deskripsi
Kode di atas berfungsi menilai siswa menggunakan `if` `else` dan `else if`. Namun kode ini di tambahkan sistem menu untuk kembali

## Unguided
### 1. Modifikasi program soal terbimbing No. 1 [grade.go] dengan menggunakan switch!
```go
package main

import "fmt"
import "bufio"
import "os"

func main() {
	var nama string
	var nilai float32
	fmt.Println("====== Program Penilaian ======")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan nama: ")
	scanner.Scan()
	nama = scanner.Text()
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	fmt.Println("====== Hasil Penilaian ======")
	fmt.Println("Nama:", nama)
	fmt.Println("Nilai:", nilai)

	switch {

		case nilai >= 90 && nilai <= 100:
			fmt.Println("Nilai A")

		case nilai >= 80 && nilai < 90:
			fmt.Println("Nilai B")

		case nilai >= 70 && nilai < 80:
			fmt.Println("Nilai C")

		case  nilai >= 60 && nilai < 70:
			fmt.Println("Nilai D")
			
		case nilai < 60:
			fmt.Println("Nilai F")
	}
}
```
##### Output
![Screenshot Output Unguided]()

#### Deskripsi
Kode di atas hanyalah modifikasi dari `grade.go` namun yang ini menggunakan switch case. secara fungsi sama

### 2. Membuat sebuah program dalam bahasa Go yang menghitung pajak penghasilan seseorang berdasarkan total penghasilannya.

```go
package main

import "fmt"

func main (){
	var uang, pajak float64
	fmt.Println("====== Masukan uang (sudah berbentuk juta) ======")
	fmt.Print("uang: " ); fmt.Scan(&uang)

	if uang <= 50 {
		pajak = (uang * 5 / 100)
		fmt.Print("pajak: ",pajak, " juta")

	} else if uang > 50 && uang <= 100 {
		pajak = (50.0 * 5.0 / 100.0) +
		((uang - 50.0) * 10.0 / 100.0)
		fmt.Print("pajak: ",pajak, " juta")

	} else if uang > 100 && uang <= 200 {
		pajak = (50.0 * 5.0 / 100.0) +
		(50.0 * 10.0 /100.0) +
		((uang - 100.0) * 15.0 / 100.0)
		fmt.Print("pajak: ",pajak, " juta")

	} else {
		pajak = (50.0 * 5.0 / 100.0) +
		(50.0 * 10.0 /100.0) +
		(100.0 * 15.0 /100.0) +
		((uang - 200.0) * 20.0 / 100.0)
		fmt.Print("pajak: ",pajak, " juta")
	}
}
```

##### Output
![Screenshot Output Unguided]()

#### Deskripsi
Kode di atas berfungsi untuk menghitung seberapa banyak pajak yang harus di bayar seseorang sesuai banyak nya uang yang di dapat dan menghitung nya di dalam kode pemrograman

## Referensi
1. *A Tour Of Go*. Diakses pada 11 Oktober 2026 melalui https://go.dev/tour/flowcontrol/9
2. *A Tour Of Go*. Diakses pada 11 Oktober 2026 melalui https://go.dev/tour/flowcontrol/5
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
