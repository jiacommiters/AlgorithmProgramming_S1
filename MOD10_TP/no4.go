package main

import "fmt"

func main() {

	var membership string
	var a, b, c, d, e int
	var discount, cashback float64
	
	fmt.Scan(&membership, &a, &b, &c, &d, &e)

	cashback = 3.1 * (float64(a) + float64(b) + float64(c))
	discount = 1.7 * (float64(c) + float64(d) + float64(e))

	if a%2 == 1 && b%2 == 1 && c%2 == 1 && d%2 == 1 && e%2 == 1 {
		cashback = 0
		discount = 1.7 * (float64(a) + float64(b) + float64(c) + float64(d) + float64(e))
	} else if a%2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0 {
		discount = 0
		cashback = 3.1 * (float64(a) + float64(b) + float64(c) + float64(d) + float64(e))
	} else if cashback > 35.00 {
		cashback = 35.00
	} else if discount > 50.00 {
		discount = 50.00
	} else if membership == "yes" {
		discount = discount * 1.15
		cashback = cashback * 1.15
	} else {

	}

	fmt.Printf("cashback %.2f dan discount %.2f\n", cashback, discount)

}
