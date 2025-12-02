package main

import "fmt"

func main() {
    var transaksi int
    var saldo int = 0
    var selesai bool = false

    for !selesai {
        fmt.Scan(&transaksi)

        if transaksi == 0 {
            selesai = true
        } else {
            saldo = saldo + transaksi
        }
    }

    fmt.Println(saldo)
}