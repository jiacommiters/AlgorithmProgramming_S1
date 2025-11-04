package main

import (
	"fmt"
)

func main() {

	var n, x, i, kartu int

	fmt.Scan(&n)
	kartu = 0

	for i = 1; i <= n; i++ {

		fmt.Scan(&x)
		kartu = kartu + (x / 5)

	}

	fmt.Print(kartu)

}
