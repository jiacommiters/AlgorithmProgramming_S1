package main

import "fmt"

func main() {

	var x byte
	fmt.Scanf("%c", &x)

	if x >= '0' && x <= '9' {

		fmt.Print("Bilangan")

	} else if x >= 'a' && x <= 'z' {

		fmt.Print("Huruf Kecil")

	} else if x >= 'A' && x <= 'Z' {

		fmt.Print("Huruf Besar")

	} else {

		fmt.Print("Simbol")

	}
}
