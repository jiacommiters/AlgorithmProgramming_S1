package main

import "fmt"

func main() {

	var alice, bob int
	var hasil bool

	fmt.Scan(&alice, &bob)

	hasil = alice == bob || alice - bob == 1 || bob - alice == 1 || bob - alice == 5 || alice - bob == 5

	fmt.Print(hasil)
}