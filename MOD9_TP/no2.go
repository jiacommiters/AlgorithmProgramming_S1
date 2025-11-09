package main
import "fmt"

func main() {

	var bus, penumpang, jumlahBus int

	fmt.Scan(&penumpang)

	bus = 45 

	jumlahBus = penumpang / bus

	if penumpang % bus >= 1 {
		jumlahBus = jumlahBus + 1
	}

	fmt.Print(jumlahBus)
}