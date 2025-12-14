package main
import "fmt"

func main() {
    var i, s int
    var v, total, totalSemua float64
    var berhenti, ada90 bool
    
    for !berhenti {
        total = 0
        ada90 = false
        for i = 0; i < 3; i++ {
            fmt.Scan(&v)
            total += vP
            totalSemua += v
            if v >= 90 {
                ada90 = true
            }
        }
        s++
        if ada90 || total >= 210 {
            berhenti = true
        }
    }  
    fmt.Printf("Total set: %d (rata-rata: %.2f WPM)", s, totalSemua/float64(s*3))
}