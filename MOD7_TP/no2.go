package main

import "fmt"

func main() {

	var d, dbd, d1, d2, d3, dsum int
	fmt.Scan(&d)

	d1 = ((d / 100) * 11) * 10000
	d2 = (((d % 100) / 10 ) * 11) * 100
	d3 = (d % 10) * 11
	dsum = d1 + d2 + d3

	dbd = dsum * dsum

	fmt.Println(dsum)
	fmt.Print(dbd)

}