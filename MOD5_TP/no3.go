package main
import "fmt"

func main() {

	var n, i, faktorial int
	fmt.Scan(&n)
	faktorial = n

	for i = n - 1; i >= 1; i-- {
		faktorial *= i

	}

	fmt.Println(faktorial)
}