package main

import "fmt"

func main() {

	type salary struct {
		nama      string
		gajiP     float64
		tunjangan float64
		potongan  float64
		totalGaji float64
	}

	var s salary

	fmt.Print("Masukan nama Pegawai : ")
	fmt.Scan(&s.nama)
	fmt.Print("Masukan Gaji Pokok Anda : ")
	fmt.Scan(&s.gajiP)
	fmt.Print("Masukan Tunjangan Anda : ")
	fmt.Scan(&s.tunjangan)
	fmt.Print("Masukan Potongan Gaji Anda : ")
	fmt.Scan(&s.potongan)

	s.totalGaji = s.gajiP + s.tunjangan - s.potongan

	fmt.Printf("Nama Pegawai\t: %s\n", s.nama)
	fmt.Printf("Gaji Pokok\t: Rp %10.2f\n", s.gajiP)
	fmt.Printf("Tunjangan\t: Rp %10.2f\n", s.tunjangan)
	fmt.Printf("Potongan\t: Rp %10.2f\n", s.potongan)
	fmt.Printf("Total Gaji\t: Rp %10.2f\n", s.totalGaji)

}
