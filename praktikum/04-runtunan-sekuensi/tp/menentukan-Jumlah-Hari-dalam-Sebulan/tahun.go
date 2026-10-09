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