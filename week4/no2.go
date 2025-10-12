package main

import "fmt"

func main() {

	type persegi struct {
		panjang  float64
		lebar    float64
		luas     float64
		keliling float64
	}

	var p persegi

	fmt.Print("Masukan Nilai Panjang : ")
	fmt.Scan(&p.panjang)

	fmt.Print("Masukan Lebar : ")
	fmt.Scan(&p.lebar)

	p.luas = p.panjang * p.lebar
	p.keliling = 2 * (p.panjang + p.lebar)

	fmt.Println("Nilai Luas Persegi Panjang : ", p.luas)
	fmt.Println("Nilai Keliling Persegi Panjang : ", p.keliling)


}