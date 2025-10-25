package main

import "fmt"

func main() {

    var s int
    var pu, su, sq, jj, tK, kPU, kSU, kSQ, kJJ, i float64
    fmt.Scan(&s)
    fmt.Scan(&pu, &su, &sq, &jj)


    for i = 1; i <= pu; i++ {
        kPU = kPU + i*0.5
    }

    for i = 1; i <= su; i++ {
        kSU = kSU + i*0.3
    }
   
    for i = 1; i <= sq; i++ {
        kSQ = kSQ + i*0.2
    }
p
    for i = 1; i <= jj; i++ {
        kJJ = kJJ + i*0.6
    }
    
    tK = (kPU + kSU + kSQ + kJJ) * float64(s)
    fmt.Printf("Total kalori terbakar hari ini sebanyak %.1f kalori\n", tK)
}