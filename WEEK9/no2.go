package main
import "fmt"

func main() {

	var harga float32
	var ujian bool
	fmt.Scan(&harga, &ujian)

	if ujian == true {

		harga = harga - (harga * 0.35)

	}

	fmt.Print(harga)

}