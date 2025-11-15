package main

import "fmt"

func main() {
    var a, b, c, d int
    fmt.Scan(&a, &b, &c, &d)
    
    var max int = a
    var min int = a
    
    if b > max {
        max = b
    }
    if c > max {
        max = c
    }
    if d > max {
        max = d
    }
    
    if b < min {
        min = b
    }
    if c < min {
        min = c
    }
    if d < min {
        min = d
    }
    
    fmt.Print(max, " ", min)
}