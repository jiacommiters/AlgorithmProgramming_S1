package main

import "fmt"

func main() {

	var  h1, h2, h3, h4, h5, c1, c2, c3, c4, c5 int
	var cek bool

	fmt.Scan(&h1, &h2, &h3, &h4, &h5)

	c1 = 20 * h1
	c2 = 20 * h2	
	c3 = 20 * h3
	c4 = 20 * h4
	c5 = 20 * h5

	cek = c1 + c2 + c3 + c4 + c5 >= 70

	fmt.Print(cek)

}