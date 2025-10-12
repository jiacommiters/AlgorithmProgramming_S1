package main

import "fmt"

func main() {

	type komposisi struct {
		tinggi float64
		berat  float64
		bmi    float64
	}

	var k komposisi

	fmt.Print("Masukan Tinggi Badan Anda : ")
	fmt.Scan(&k.tinggi)

	fmt.Print("Masukan Berat Badan Anda : ")
	fmt.Scan(&k.berat)

	// if k.tinggi > 2.0 {
	// 	k.tinggi / 100
	// } else {
		
	// }
		
	k.bmi = k.berat / k.tinggi
	fmt.Print("BMI Anda : ", k.bmi)

}