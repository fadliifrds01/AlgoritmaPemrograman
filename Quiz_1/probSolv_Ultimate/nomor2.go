package main

import "fmt"

func main() {
	var angka int

	fmt.Println("Masukkan bilangan tiga digit :")
	fmt.Scan(&angka)

	ratusan := angka / 100
	puluhan := (angka / 10) % 10
	satuan := angka % 10

	hasil := satuan*100 + puluhan*10 + ratusan

	fmt.Println("Hasil :", hasil)
}

// Program TukarDigit

// Kamus Data
//     Angka : integer
//     Ratusan : integer
//     Puluhan : integer
//     Satuan : integer
//     Hasil : integer

//     INPUT Angka

//     Ratusan <- Angka / 100
//     Puluhan <- (Angka / 10) % 10
//     Satuan <- Angka % 10

//     Hasil <- Satuan * 100 + Puluhan * 10 + Ratusan

//     OUTPUT Hasil

// EndProgram
