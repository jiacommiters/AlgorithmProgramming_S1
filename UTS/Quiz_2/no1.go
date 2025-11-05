package main

import "fmt"

func main() {

	var angka, i, n, n2, nilaiAngka int

	fmt.Scan(&angka)
	fmt.Scan(&n, &n2)

	nilaiAngka = 1

	for i = 1; i <= n; i++ {

		nilaiAngka = nilaiAngka * angka
		fmt.Println(nilaiAngka)

	}

	for i = 1; i <= n2-n; i++ {

		nilaiAngka = nilaiAngka * angka
		fmt.Println(nilaiAngka)

	}

}
