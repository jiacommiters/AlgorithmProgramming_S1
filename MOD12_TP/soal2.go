package main 
import "fmt"

func main() {

	var batas, x, hasil int
	fmt.Scanln(&batas)

	for hasil <= batas {

		fmt.Scan(&x)
		hasil = hasil + x
		
		
	}
``
	fmt.Print(hasil)

}