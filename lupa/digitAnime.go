package main
import "fmt"

func main() {

	var x, id1, id2, id3 int
	fmt.Print("Masukan nilai x untuk dilakukan operasi mod : ")
	fmt.Scan(&x)

	id1 = x % 1000 / 100
	id2 = x % 100 / 10
	id3 = x % 10

	fmt.Print(id1, id2, id3)

}