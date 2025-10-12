package main
import "fmt"

func main() {
    type saham struct {
        beli     float64
        jual     float64
        jSaham   int
        tinvest  float64
        tjual    float64
        kotor    float64
        biaya    float64
        pajak    float64
        bersih   float64
    }

    var s saham
    
    fmt.Print("Harga beli saham per lembar: ")
    fmt.Scan(&s.beli)
    fmt.Print("Harga jual saham per lembar: ")
    fmt.Scan(&s.jual)
    fmt.Print("Jumlah saham: ")
    fmt.Scan(&s.jSaham)

    s.tinvest = s.beli * float64(s.jSaham)
    s.tjual = s.jual * float64(s.jSaham)
    s.kotor = s.tjual - s.tinvest
    s.biaya = s.tjual * 0.002

    if s.kotor > 0 {
        s.pajak = s.kotor * 0.1
    } else {
        s.pajak = 0
    }
    s.bersih = s.kotor - s.biaya - s.pajak

    
    fmt.Println("\n=== LAPORAN TRANSAKSI SAHAM ===")
    fmt.Printf("Harga Beli per Saham\t: Rp %.2f\n", s.beli)
    fmt.Printf("Harga Jual per Saham\t: Rp %.2f\n", s.jual)
    fmt.Printf("Jumlah Saham\t\t: %d lembar\n", s.jSaham)
    fmt.Printf("Total Investasi\t\t: Rp %.2f\n", s.tinvest)
    fmt.Printf("Total Penjualan\t\t: Rp %.2f\n", s.tjual)
    fmt.Printf("Laba/Rugi Kotor\t\t: Rp %.2f\n", s.kotor)
    fmt.Printf("Biaya Transaksi (0.2%%)\t: Rp %.2f\n", s.biaya)
    fmt.Printf("Pajak (10%%)\t\t: Rp %.2f\n", s.pajak)
    fmt.Println("---------------------------------")
    fmt.Printf("Laba/Rugi Bersih\t: Rp %.2f\n", s.bersih)
}