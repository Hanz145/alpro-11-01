package main
import "fmt"

func main(){
	//var
	var kelembab, bridge, jml int
	terendah := 999
	tertinggi := -999
	var loop int = 0
	var rata2 float32

	//percabangan
	for {
		fmt.Scan(&kelembab)
		if kelembab < terendah {
			terendah = kelembab
		}
		if kelembab > tertinggi{
			tertinggi = kelembab
		}
		if kelembab == bridge {
			break
		}
		loop ++
		jml = jml + kelembab
		rata2 = float32(jml)/float32(loop)
		bridge = kelembab
	}
	fmt.Println(tertinggi, terendah, rata2)
}