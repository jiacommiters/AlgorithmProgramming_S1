package main

import "fmt"

func main() {

	var x, n, i, vn int
	var s float64

	fmt.Scan(&x, &n)

	vn = n + 1
	for i = 1; i <= n; i++ {


		s = s + (float64(vn) - float64(i)) / (float64(x) * float64(i))
		

	}

	fmt.Printf("%.3f", s)
}