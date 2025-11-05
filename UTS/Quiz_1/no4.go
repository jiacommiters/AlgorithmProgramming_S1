package main

import "fmt"

func main() {

	var y, a, b, c, d, e, f float64

	fmt.Scan(&a, &b, &c, &d, &e, &f)

	y = (((a - b) * (c - d)) / e) + f

	fmt.Print(y)

}