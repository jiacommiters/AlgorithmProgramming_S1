package main

import "fmt"

func main() {

	var panjang, itteration, angka, perkalian int
	fmt.Scan(&panjang)

	for itteration = 0; itteration < panjang; itteration++ {
		fmt.Scan(&angka)
		perkalian = itteration + 1
		angka = angka * perkalian
		fmt.Print(angka, "")

	}

}
