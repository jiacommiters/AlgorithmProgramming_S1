package main
import "fmt"

func main() {

	var hasil, x, y, n, i int
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&x, &y)
		hasil = x * y
		fmt.Println(hasil)
	}

}