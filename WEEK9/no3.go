package main
import "fmt"

func main() {
    var x string
    fmt.Scan(&x)
    
    if x == "A" || x == "I" || x == "U" || x == "E" || x == "O" ||
       x == "a" || x == "i" || x == "u" || x == "e" || x == "o" {
        fmt.Print("bukan konsonan")
    } else if (x >= "A" && x <= "Z") || (x >= "a" && x <= "z") {
        fmt.Print("konsonan")
    } else {
        fmt.Print("bukan konsonan")
    }
}p