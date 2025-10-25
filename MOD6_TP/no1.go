package main

import "fmt"

func main() {

	var a, b, i, total int
	fmt.Scan(&a, &b)

	total = 0
	for i = 1; i <= b; i++ {
		total += a * i
	}

	fmt.Println(total)


}
