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
