package main
import "fmt"

func main() {
    var buku int
    var diskon, total float64
    var member rune
    fmt.Scanf("%d %c", &buku, &member)
    switch member {
    case 'A':
        if buku < 5 {
            diskon = 0.1
        } else if buku > 10 {
            diskon = 0.3
        } else {
            diskon = 0.2
        }
    case 'B':
        if buku < 5 {
            diskon = 0.05
        } else if buku > 10 {
            diskon = 0.15
        } else {
            diskon = 0.1
        }
    case 'C':
        if buku > 10 {
            diskon = 0.1
        } else {
            diskon = 0.05
        }
    case 'N':
        if buku > 10 {
            diskon = 0.05
        } else {
            diskon = 0.0
        }
    default:
        diskon = 0.0
    }
    total = float64(buku*10000) - float64(buku*10000)*diskon
    fmt.Print("Rp ", total)
}