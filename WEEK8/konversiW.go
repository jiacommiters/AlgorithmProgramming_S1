package main

import "fmt"

func main() {
	// Data siswa
	students := []struct {
		no     int
		name   string
		gender string
	}{
		{1, "MUHAMMAD DZIA ILHAQI", "P"},
		{2, "MAFTUKH ASHFANDI", "P"},
		{3, "KJELL MIKHAIL RAGNVALD", "P"},
		{4, "AISHA HAYA AZ-ZAHRA KHOIRUNNISA", "W"},
		{5, "ADRA ATHALLAH OCTAVIADI", "P"},
		{6, "RADITYA RIANDIKA IRFANI", "P"},
		{7, "MADTROPE'I PUTRA ANDOKI", "P"},
		{8, "AHSANUL HADI FAUZI", "P"},
		{9, "RAFAEL PUTRA ANANTA SARAGIH", "P"},
		{10, "RAFI AHMAD MUZZAKI", "P"},
		{11, "FAIZ RAMDLAN ANSHORI", "P"},
		{12, "MUH ALIF FAYYAZ", "P"},
		{13, "FAHRY AFIF SIREGAR", "P"},
		{14, "RAJWA NAILA ARKAN", "W"},
		{15, "GRACE MARTINA CAROLINE ARITONANG", "W"},
		{16, "ALIYA AMMARANISA KURNIAWAN", "W"},
		{17, "ALLYSA MUTIARA PUTRI NABABAN", "W"},
		{18, "MUHAMMAD FAUZAN", "P"},
		{19, "KEVIN ADITYA PRATAMA", "P"},
		{20, "M. FAJAR RAMADHAN", "P"},
		{21, "ARSHANTA AYU KUSUMANINGTYAS", "W"},
		{22, "DYLAN SETYAWAN EKAPRAJA", "P"},
		{23, "DENI KURNIAWAN", "P"},
		{24, "DIVANCA JADZIA HASAN", "W"},
		{25, "RASHAD OVIN ASTANTO", "P"},
		{26, "ENDRA", "P"},
		{27, "MUHAMMAD TEGUH TRIYANUGRAHA", "P"},
		{28, "DIEGO FORTUNA", "P"},
		{29, "MUHAMMAD BINTANG NABEEL", "P"},
		{30, "MUHAMMAD RAFI HAZIMULFIKRI", "P"},
		{31, "ATTALA DEVAL ALRYAN", "P"},
		{32, "ARDHIAN GALIH PRATAMA", "P"},
		{33, "MICHAEL VINDENOVAN LIMBONG", "P"},
		{34, "AHMAD SYADALI BIMA BIRAWA", "P"},
		{35, "MADE BAYU ADNYANA", "P"},
		{36, "ARIA CHRISNA BAYU", "P"},
		{37, "MUHAMAD NUR MAOLANA AMALIA", "P"},
		{38, "MUH. SYAH DEVAN AL GHIFARI", "P"},
		{39, "CAYLA DIANDRA PUTRI SILITONGA", "W"},
	}

	kelompok1 := []string{
		students[38].name, // CAYLA DIANDRA PUTRI SILITONGA (W)
		students[24].name, // RASHAD OVIN ASTANTO (P)
		students[25].name, // ENDRA (P)
	}

	// Sisanya dibagi secara manual dengan memperhatikan aturan wanita
	kelompok2 := []string{
		students[3].name, // AISHA HAYA AZ-ZAHRA KHOIRUNNISA (W)
		students[1].name, // MAFTUKH ASHFANDI (P)
		students[4].name, // ADRA ATHALLAH OCTAVIADI (P)
	}

	kelompok3 := []string{
		students[13].name, // RAJWA NAILA ARKAN (W)
		students[6].name,  // MADTROPE'I PUTRA ANDOKI (P)
		students[7].name,  // AHSANUL HADI FAUZI (P)
	}

	kelompok4 := []string{
		students[14].name, // GRACE MARTINA CAROLINE ARITONANG (W)
		students[8].name,  // RAFAEL PUTRA ANANTA SARAGIH (P)
		students[9].name,  // RAFI AHMAD MUZZAKI (P)
	}

	kelompok5 := []string{
		students[15].name, // ALIYA AMMARANISA KURNIAWAN (W)
		students[10].name, // FAIZ RAMDLAN ANSHORI (P)
		students[11].name, // MUH ALIF FAYYAZ (P)
	}

	kelompok6 := []string{
		students[16].name, // ALLYSA MUTIARA PUTRI NABABAN (W)
		students[12].name, // FAHRY AFIF SIREGAR (P)
		students[17].name, // MUHAMMAD FAUZAN (P)
	}

	kelompok7 := []string{
		students[20].name, // ARSHANTA AYU KUSUMANINGTYAS (W)
		students[18].name, // KEVIN ADITYA PRATAMA (P)
		students[19].name, // M. FAJAR RAMADHAN (P)
	}

	kelompok8 := []string{
		students[23].name, // DIVANCA JADZIA HASAN (W)
		students[21].name, // DYLAN SETYAWAN EKAPRAJA (P)
		students[22].name, // DENI KURNIAWAN (P)
	}

	kelompok9 := []string{
		students[0].name, // MUHAMMAD DZIA ILHAQI
		students[2].name, // KJELL MIKHAIL RAGNVALD
		students[5].name, // RADITYA RIANDIKA IRFANI
	}

	// Kelompok tanpa wanita
	kelompok10 := []string{
		students[26].name, // MUHAMMAD TEGUH TRIYANUGRAHA
		students[27].name, // DIEGO FORTUNA
		students[28].name, // MUHAMMAD BINTANG NABEEL
	}

	kelompok11 := []string{
		students[29].name, // MUHAMMAD RAFI HAZIMULFIKRI
		students[30].name, // ATTALA DEVAL ALRYAN
		students[31].name, // ARDHIAN GALIH PRATAMA
	}

	kelompok12 := []string{
		students[32].name, // MICHAEL VINDENOVAN LIMBONG
		students[33].name, // AHMAD SYADALI BIMA BIRAWA
		students[34].name, // MADE BAYU ADNYANA
	}

	kelompok13 := []string{
		students[35].name, // ARIA CHRISNA BAYU
		students[36].name, // MUHAMAD NUR MAOLANA AMALIA
		students[37].name, // MUH. SYAH DEVAN AL GHIFARI
	}

	// Menampilkan semua kelompok
	kelompok := [][]string{
		kelompok1, kelompok2, kelompok3, kelompok4, kelompok5,
		kelompok6, kelompok7, kelompok8, kelompok9, kelompok10,
		kelompok11, kelompok12, kelompok13,
	}

	for i, k := range kelompok {
		fmt.Printf("Kelompok %d:\n", i+1)
		for j, anggota := range k {
			fmt.Printf("  %d. %s\n", j+1, anggota)
		}
		fmt.Println()
	}
}
