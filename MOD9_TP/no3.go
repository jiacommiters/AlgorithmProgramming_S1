package main
import "fmt"


func main() {

	var Halice, Hbob, usd, Palice, Pbob int
	fmt.Scan(&Halice, &Hbob)
	usd = 2000

	Palice = Halice / 2000
	Pbob = Hbob / 2000

	if Halice % usd > 0 {
		Palice = Palice + 1
	}
	if Hbob % usd > 0 {
		Pbob = Pbob + 1
	}

	fmt.Println("PT. Alice memerlukan", Palice, "lembar USD 2,000")
	fmt.Println("PT. Bob memerlukan", Pbob, "lembar USD 2,000")



}