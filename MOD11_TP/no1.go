package main

import "fmt"

func main() {

	var x int
	fmt.Scan(&x)

	switch x {

	case 1:
		fmt.Println("Januari")
	case 2:
		fmt.Println("Februari")
	case 3:
		fmt.Println("Maret")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("Mei")
	case 6:
		fmt.Println("Juni")
	case 7: 
		fmt.Println("July")
	case 8: 
		fmt.Println("Agustus")
	case 9:
		fmt.Println("September")
	case 10: 
		fmt.Println("Oktober")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("Dececmber")
	default:
		fmt.Println("Input is Invalid")
	}
}
