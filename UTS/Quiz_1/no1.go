package main

import "fmt"

func main() {

	var j, m, d int
	var kj, km, kd, hwaktu bool

	fmt.Scan(&j, &m, &d)

	kj = j > 0 && j <= 12
	km = m >= 0 && m < 60
	kd = d >= 0 && d < 60

	hwaktu = kj && km && kd


	fmt.Print(hwaktu)

}
