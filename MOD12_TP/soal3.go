package main

import "fmt"

func main() {

	var text byte
	var jumlah int

	fmt.Scanf("%c", &text)

	// for text != '#' {
	//     jumlah = jumlah + 1
	//     fmt.Scanf("%c", &text)
	// }

	for ((text >= 'A' && text <= 'Z') || (text >= 'a' && text <= 'z') || (text >= '0' && text <= '9' || text == ' ')) && text != '#' {

		jumlah = jumlah + 1
		fmt.Scanf("%c", &text)

	}

	// for (!(text >= 'A' && text <= 'Z') || !(text >= 'a' && text <= 'z') || !(text >= '0' && text <= '9')) && text == ' ' && text != '#' {
	// 	jumlah2 = jumlah2 + 1
	// 	fmt.Scanf("%c", &text)
	// }

	fmt.Print(jumlah)

}
