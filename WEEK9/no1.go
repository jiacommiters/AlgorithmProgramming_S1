package main
import "fmt"

func main() {

	var mahasiswa, mobil, jumlahM int
	fmt.Scan(&mahasiswa)

	mobil = 7
	jumlahM = mahasiswa / mobil

	if mahasiswa % mobil > 0 {

		jumlahM = jumlahM + 1

	}

	fmt.Println(jumlahM)


}