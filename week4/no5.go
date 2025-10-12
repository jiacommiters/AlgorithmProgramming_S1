package main

import "fmt"

func main() {
	var hBeli, hJual, jSaham, tInvest, tPenjualan, kKotor, bTransaksi, pKeuntungan, kBersih float64

	fmt.Print("Harga Beli per Saham (Rp): ")
	fmt.Scan(&hBeli)
	fmt.Print("Harga Jual per Saham (Rp): ")
	fmt.Scan(&hJual)
	fmt.Print("Jumlah Saham (lembar): ")
	fmt.Scan(&jSaham)
	fmt.Println()

	tInvest = hBeli * jSaham
	tPenjualan = hJual * jSaham
	kKotor = tPenjualan - tInvest
	bTransaksi = tPenjualan * (0.2 / 100)
	pKeuntungan = 0.0
	pKeuntungan = (map[bool]float64{true: kKotor * 0.1})[kKotor > 0]
	kBersih = kKotor - bTransaksi - pKeuntungan

	fmt.Printf("Harga Beli per Saham\t: Rp %.2f\n", hBeli)
	fmt.Printf("Harga Jual per Saham\t: Rp %.2f\n", hJual)
	fmt.Printf("Jumlah Saham\t\t: %.0f lembar\n", jSaham)
	fmt.Printf("Total Investasi\t\t: Rp %.2f\n", tInvest)
	fmt.Printf("Total Penjualan\t\t: Rp %.2f\n", tPenjualan)
	fmt.Printf("Keuntungan Kotor\t: Rp %.2f\n", kKotor)
	fmt.Printf("Biaya Transaksi\t\t: Rp %.2f\n", bTransaksi)
	fmt.Printf("Pajak Keuntungan\t: Rp %.2f\n", pKeuntungan)
	fmt.Printf("Keuntungan Bersih\t: Rp %.2f\n", kBersih)
}
