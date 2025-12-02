package main

import "fmt"

func main() {
    var bilangan int
    var total int = 0
    var digit int

    fmt.Scan(&bilangan)

    for bilangan > 0 {
        digit = bilangan % 10
        fmt.Print(digit, " ")
        total = total + digit
        bilangan = bilangan / 10
    }

    fmt.Println()
    fmt.Println(total)
}