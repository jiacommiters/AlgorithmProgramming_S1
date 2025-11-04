package main

import "fmt"

func main() {

	var x, n, i int
	var y float64
	y = 1
	n = 10
	fmt.Scan(&x)

	for i = 1; i <= n; i++ {

		y = ((y + (float64(x) / y)) / 2)

	}

	fmt.Printf("%.8f", y)
}
