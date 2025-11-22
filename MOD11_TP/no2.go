package main
import "fmt"

func main() {

	var angka1, angka2, operasi, hasil int
	fmt.Scan(
		&angka1,
		&angka2,
	)

	fmt.Scan(&operasi)

	switch operasi {

	case 1:
		hasil = angka1 + angka2
		fmt.Print(hasil)
	case 2: 
		hasil = angka1 - angka2
		fmt.Print(hasil)
	case 3: 
		hasil = angka1 * angka2
		fmt.Print(hasil)
	case 4: 
		if angka2 == 0{
			fmt.Print("Tidak Terdefinisi")
		} else {
			hasil = angka1 / angka2
			fmt.Print(hasil)
		}
	case 5:
		if angka2 == 0{
			fmt.Print("Tidak Terdefinisi")
		} else {
			hasil = angka1 % angka2
			fmt.Print(hasil)
		}
	case 6:
		if angka2 == 0{
			fmt.Print("Tidak Terdefinisi")
		} else {
			hasil = angka1 / angka2
			fmt.Printf("%f", hasil)
		}
	}

}