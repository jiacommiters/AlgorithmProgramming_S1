package main
import "fmt"

func main() {

	var a1, a2, a3 int
	var hasil bool
	fmt.Scan(&a1, &a2, &a3)

	hasil = (a1 > a2 && a2 > a3) || (a1 < a2 && a2 < a3)
	fmt.Print(hasil)

}