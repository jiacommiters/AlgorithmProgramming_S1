package main
import "fmt"

func main(){

	var i, n int
	fmt.Scan(&n)

	fmt.Print(n)
	for i = n - 1; i >= 1; i-- {
		fmt.Print(" x ", i)
	}

	

}
