package main

import "fmt"

func main() {

	var jumlah, harga, total float64
	var typeMember string
	fmt.Scan(&jumlah, &typeMember)

	harga = 10000

	switch typeMember {

	case "A":

		switch {
		case jumlah > 0 && jumlah < 5:
			total = total + ((harga * jumlah) - ((harga * jumlah) * 0.1))
		case jumlah >= 5 && jumlah <= 10:
			total = total + ((harga * jumlah) - ((harga * jumlah) * 0.2))
		case jumlah > 10:
			total = total + ((harga * jumlah) - ((harga * jumlah) * 0.3))
		}

	case "B":

		switch {
		case jumlah > 0 && jumlah < 5:
			total = total + ((harga * jumlah) - ((harga * jumlah) * 0.05))
		case jumlah >= 5 && jumlah <= 10:
			total = total + ((harga * jumlah) - ((harga * jumlah) * 0.1))
		case jumlah > 10:
			total = total + ((harga * jumlah) - ((harga * jumlah) * 0.15))
		}

	case "C":
		switch {
		case jumlah > 0 && jumlah < 5:
			total = total + (harga * jumlah)
		case jumlah >= 5 && jumlah <= 10:
			total = total + ((harga * jumlah) - ((harga * jumlah) * 0.05))
		case jumlah > 10:
			total = total + ((harga * jumlah) - ((harga * jumlah) * 0.1))
		}

	case "N":
		switch {
		case jumlah > 0 && jumlah < 5:
			total = total + (harga * jumlah)
		case jumlah >= 5 && jumlah <= 10:
			total = total + (harga * jumlah)
		case jumlah > 10:
			total = total + ((harga * jumlah) - ((harga * jumlah) * 0.05))
		}
	}

	fmt.Println("Rp", total)

}
