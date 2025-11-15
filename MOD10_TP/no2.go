package main

import "fmt"

func main() {

	var x int
	fmt.Scan(&x)

	if x >= 0 && x <= 12 {
		fmt.Print("Anak-Anak")
	} else if x >= 13 && x <= 17 {
		fmt.Print("Remaja")
	} else if x >= 18 && x <= 55 {
		fmt.Print("Dewasa")
	} else if x > 55 {
		fmt.Print("Lansia")
	} else{
		
	}

}