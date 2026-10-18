package main

import "fmt"

func main() {
	var kembalian int

	fmt.Println("Masukkan jumlah kembalian :")
	fmt.Scan(&kembalian)

	sepuluhRibu := kembalian / 10000
	sisa := kembalian % 10000

	limaRibu := sisa / 5000
	sisa = sisa % 5000

	seribu := sisa / 1000

	fmt.Println(sepuluhRibu)
	fmt.Println(limaRibu)
	fmt.Println(seribu)
}

// Program KasirAlgomart

// Kamus Data
//     Kembalian : integer
//     SepuluhRibu : integer
//     Sisa : integer
//     LimaRibu : integer
//     Seribu : integer

//     INPUT Kembalian

//     SepuluhRibu <- Kembalian / 10000
//     Sisa <- Kembalian % 10000

//     LimaRibu <- Sisa / 5000
//     Sisa <- Sisa % 5000

//     Seribu <- Sisa / 1000

//     OUTPUT SepuluhRibu
//     OUTPUT LimaRibu
//     OUTPUT Seribu

// EndProgram
