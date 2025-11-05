package main

import "fmt"

func main() {

	var h1, h2, h3, h4, h5 int
	var jenny, c1, c2, c3, c4, c5, rate float64

	fmt.Scan(&h1, &h2, &h3, &h4, &h5)

	jenny = 111.33

	c1 = jenny * float64(h1)
	c2 = jenny * float64(h2)
	c3 = jenny * float64(h3)
	c4 = jenny * float64(h4)
	c5 = jenny * float64(h5)

	rate = (c1 + c2 + c3 + c4 + c5) / 5

	fmt.Print(rate)

}
