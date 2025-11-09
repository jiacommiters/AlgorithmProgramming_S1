package main
import "fmt"

func main() {

	// var x, y, z, sX, sY, sZ int
	// fmt.Scan(&x, &y, &z)

	// if x < 0 {
	// 	sX = (x * -2) / 2 
	// }

	// if x > 0 {
	// 	sX = x * 1
	// }

	// if y < 0 {
	// 	sY = (y * -2) / 2
	// }

	// if y > 0 {
	// 	sY = y * 1
	// }

	// if z < 0 {
	// 	sZ = (z * -2) / 2
	// } 

	// if z > 0 {
	// 	sZ = z * 1
	// } 

	// fmt.Print(sX, sY, sZ)

	var  i, angka int
	
	for i = 1; i <= 3; i++ {

		fmt.Scan(&angka)
		
		if angka < 0 {
			angka = -angka
		}

		fmt.Println(angka)

	}

}