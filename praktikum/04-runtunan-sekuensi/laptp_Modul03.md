# <h1 align="center">Tugas Pendahuluan Modul 04 - runtunan sekuensi</h1>
<p align="center">Reihan - 109092600002</p>

### 1. Evaluasi Ekspresi Kontrol
```go
package main

import "fmt"

func main(){
	intNum := 5
	if  !(intNum > 3) || intNum <= 5  {
		fmt.Println("True")
	} else {
		fmt.Print("False")
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![alt text](https://github.com/Hanz145/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/tp/Evaluasi-Ekspresi-Kontrol-dalam-Go/output.png?raw=true)


#### Deskripsi
Mengecek apakah variabel `intNum` tidak sama dengan lebih dari tiga atau kurang dari samadengan lima.

### 2. Switch Case

```go
package main

import "fmt"

func main() {
	fmt.Println("====== Switch Case ======")
	fmt.Println("Masukan bilangan dan lihat apakah bilangan tersebut > 10, <= 25, atau == 10 : ")
	fmt.Print("Masukkan bilangan: ")
	var bilangan int
	fmt.Scan(&bilangan)

	switch {
	case bilangan > 10:
		fmt.Println("Bilangan lebih besar dari 10")
	case bilangan <= 25:
		fmt.Println("Bilangan lebih kecil atau sama dengan 25")
	case bilangan == 10:
		fmt.Println("Bilangan sama dengan 10")
	default:
		fmt.Println("Bilangan tidak memenuhi kondisi")
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/Hanz145/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/tp/Switch-Case/outout.png?raw=true)


#### Deskripsi
Dari kode di atas kita bisa tau apakah bilangan yang kita masukan Lebih dari sepuluh, kurang dari sama dengan 25, atau == 10

### 3. Tracing Evaluasi Pernyataan Kondisi

```go
package main
import "fmt"  
func main() { 
	x := 10 
	y := 5 
	z := 15 
	result := 0
	if x > 5 {	
		if y < 10 {
			result = x + y 
		} else {
			result = x - y
		}
	}
	if z > 10 && x == 10 {
		result += z
	} else {
		result = z - x 
	}
	if x == 10 || y > 10 {
		result += 5
	} else if y == 5 && z > 10 {
		result -= 5
	} else {
		result *= 2
	}
	if !(x < 15 && y < 10) {
		result += 10
	} else {
		result -= 10
	}
	fmt.Println("Nilai akhir result:", result)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/Hanz145/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/tp/Tracing-Evaluasi_Pernyataan_Kondisi/output.png?raw=true)

#### Deskripsi
Kode di atas menampilkan hasil sesuai kondisi yang terpenuhi menggunakan `if` dan `else`

### 4. menentukan Jumlah Hari dalam Sebulan

```go
package main


import "fmt"


func main (){
	var tiga0 int = 30
	var tiga1 int = 31
	var feb int = 28
	var kab int = 29
	var tahun int
	var bulan string
	fmt.Println("====== Masukan Tahun dan bulan ======")
	fmt.Print("tahun: "); fmt.Scan(&tahun)
	fmt.Print("bulan: "); fmt.Scan(&bulan)
	if tahun % 4 == 0 {
		if bulan == "Jan" || bulan == "Mar" || bulan == "Mei" || bulan == "Jul"|| bulan == "Aug" || bulan == "Okt" || bulan == "Des" {
			fmt.Print("hari: ",tiga1) 
		} else if bulan == "Apr" || bulan == "Jun" || bulan == "Sep" || bulan == "Nov"  {
			fmt.Print("hari: ",tiga0) 
		} else if bulan == "Veb" {
			fmt.Print("hari: ",kab)
		}
	} else {
		if bulan == "Jan" || bulan == "Mar" || bulan == "Mei" || bulan == "Jul"|| bulan == "Aug" || bulan == "Okt" || bulan == "Des" {
			fmt.Print("hari: ",tiga1) 
		} else if bulan == "Apr" || bulan == "Jun" || bulan == "Sep" || bulan == "Nov"  {
			fmt.Print("hari: ",tiga0) 
		} else if bulan == "Veb" {
			fmt.Print("hari: ",feb)
		}
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/Hanz145/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/tp/menentukan-Jumlah-Hari-dalam-Sebulan/output.png?raw=true)

#### Deskripsi
Kode di atas menampilkan berapa banyak hari sesuai bulan dan tahun kabisat dengan akurat

## Kesimpulan
tujuan praktikum bertujuan untuk mengetahui cara menggunakan `if`, `else`, dan `switch case`