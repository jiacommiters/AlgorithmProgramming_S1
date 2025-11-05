package main

import "fmt"

func main() {

	var angka, i, n, n2, nilaiAngka, j int

	fmt.Scan(&angka)
	fmt.Scan(&n, &n2)

	for i = n; i <= n2; i++ {

		nilaiAngka = 1

		for j = 1; j <= i; j++ {
			nilaiAngka = nilaiAngka * angka
			fmt.Print(nilaiAngka)

		}

	}

}
