package main
import "fmt"

func main() {

	var harga float32
	var ujian bool
	fmt.Scan(&harga, &ujian)

	if ujian == true {

		harga = harga * 0.65

	}

	fmt.Print(harga)

}