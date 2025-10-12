package main
import "fmt"

func main() {

	var avtur, pound, jumlahKg, jumlahL, volume float64
	fmt.Print("Masukan jumlah avtur : ")
	fmt.Scan(&avtur)

	pound = 0.45359237
	volume = 0.80

	jumlahKg = avtur * pound
	jumlahL = jumlahKg / volume

	fmt.Printf("%.6f kg\n", jumlahKg)
	fmt.Printf("%.6f L\n", jumlahL)

}