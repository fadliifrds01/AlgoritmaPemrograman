package main

import "fmt"

func main() {
	var angka, sisipan int

	fmt.Println("Masukkan Ankga :")
	fmt.Scan(&angka, &sisipan)

	duaDigitDepan := angka / 100
	duaDigitBelakang := angka % 100
	hasil := duaDigitDepan*1000 + sisipan*100 + duaDigitBelakang

	fmt.Println(hasil)
}

// Program SisipkanAngka

// Kamus Data
//     Angka : integer
//     Sisipan : integer
//     DuaDigitDepan : integer
//     DuaDigitBelakang : integer
//     Hasil : integer

//     INPUT Angka
//     INPUT Sisipan

//     DuaDigitDepan <- Angka / 100
//     DuaDigitBelakang <- Angka % 100
//     Hasil <- DuaDigitDepan * 1000 + Sisipan * 100 + DuaDigitBelakang

//     OUTPUT Hasil

// EndProgram
