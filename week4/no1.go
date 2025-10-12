package main
import "fmt"

func main() {

	type transaksi struct {
		nama_barang string
		jumlah int
		harga int
		total_harga int

	}

	var t transaksi

	fmt.Print("Masukan Nama Barang : ")
	fmt.Scan(&t.nama_barang)

	fmt.Print("Masukan Jumlah Barang : ")
	fmt.Scan(&t.jumlah)
	
	fmt.Print("Masukan Harga Barang : ")
	fmt.Scan(&t.harga)

	t.total_harga = t.harga * t.jumlah
	fmt.Println("Barang Yang Anda Beli : ", t.nama_barang)
	fmt.Println("Total Transaksi Anda adalah : ", t.total_harga)


}