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
![alt text](image.png)


#### Deskripsi
Mengecek apakah variabel `intNum` tidak sama dengan lebih dari tiga atau kurang dari samadengan lima.

### 2. Boolean

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
![Screenshot Output Unguided](https://github.com/Hanz145/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/tp/bool/output.png?raw=true)


#### Deskripsi
Dari kode di atas kita bisa tau apakah bilangan yang kita masukan Lebih dari sepuluh, kurang dari sama dengan 25, atau == 10

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
![Screenshot Output Unguided](image-1.png)

## Kesimpulan
tujuan praktikum bertujuan untuk mengetahui cara mengkonverensi, menemukan hasil bagi, dan menentukan true dan false di dalam go lang