package main

import "fmt"

func main() {
	var angka int

	fmt.Println("Masukkan Angka :")
	fmt.Scan(&angka)

	ribuan := angka / 1000
	ratusan := (angka / 100) % 10
	puluhan := (angka / 10) % 10
	satuan := angka % 10

	hasil := ribuan < ratusan && ratusan < puluhan && puluhan < satuan

	fmt.Println(hasil)
}

// PROGRAM DigitMenaik
// KAMUS DATA
//     angka : integer
//     ribuan, ratusan, puluhan, satuan : integer
//     hasil : boolean

// ALGORITMA
//     input(angka)

//     ribuan <- angka div 1000
//     ratusan <- (angka div 100) mod 10
//     puluhan <- (angka div 10) mod 10
//     satuan <- angka mod 10

//     hasil <- (ribuan < ratusan) AND
//              (ratusan < puluhan) AND
//              (puluhan < satuan)

//     output(hasil)
// ENDPROGRAM