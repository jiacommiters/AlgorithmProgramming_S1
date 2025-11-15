package main
import "fmt"

func main() {
    var x string
    fmt.Scan(&x)
    
    if x == "A" || x == "I" || x == "U" || x == "E" || x == "O" ||p
       x == "a" || x == "i" || x == "u" || x == "e" || x == "o" {
        fmt.Print("bukan konsonan")
    } else {
        fmt.Print("konsosnan")
    }
}