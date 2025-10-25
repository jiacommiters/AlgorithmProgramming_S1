package main

import "fmt"

func main() {

	var panjang, iteration, angka, operasi1, operasi2, hasil int
	fmt.Scan(&panjang)

	for iteration = 0; iteration < panjang; iteration++ {

		fmt.Scan(&angka)
		operasi1 = angka / 1000
		operasi2 = angka % 10
		hasil = hasil + operasi1 + operasi2
		fmt.Print(hasil, " ")

	}

}
