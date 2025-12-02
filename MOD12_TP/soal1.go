package main

import "fmt"

func main() {
    var angka string
    var jumlah int
    var i int
    var digit int
    var genap int
    
    fmt.Scan(&angka)
    
    jumlah = 0
    i = 0
    
    for i < len(angka) {
        digit = int(angka[i] - '0')
        genap = 1 - (digit % 2)
        jumlah = jumlah + genap
        i = i + 1
    }
    
    fmt.Println(jumlah)
}