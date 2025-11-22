package main

import "fmt"

func main() {
	var membership string
	var a, b, c, d, e int
	var discount, cashback float64
	var totalAll float64

	fmt.Scan(&membership, &a, &b, &c, &d, &e)

	totalAll = float64(a + b + c + d + e)

	if a%2 != 0 && b%2 != 0 && c%2 != 0 && d%2 != 0 && e%2 != 0 {
		discount = 1.7 * totalAll
		cashback = 0
	} else if a%2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0 {
		discount = 0
		cashback = 3.1 * totalAll
	} else {
		cashback = 3.1 * float64(a+b+c)
		discount = 1.7 * float64(c+d+e)
	}

	if membership == "yes" {
		discount = discount * 1.15
		cashback = cashback * 1.15
	}

	if cashback > 35 {
		cashback = 35
	}
	if discount > 50 {
		discount = 50
	}

	fmt.Printf("Cashback: %.2f & Discount: %.2f\n", cashback, discount)
}
