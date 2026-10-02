package main

import "fmt"

func main() {
	var angka int32

	fmt.Scan(&angka)

	tanda := angka >> 31
	hasil := (angka ^ tanda) - tanda

	fmt.Println(hasil)
}

// Program NilaiMutlak

// Kamus Data
//     Angka : integer
//     Tanda : integer
//     Hasil : integer

//     INPUT Angka

//     Tanda <- Angka >> 31
//     Hasil <- (Angka XOR Tanda) - Tanda

//     OUTPUT Hasil

// EndProgram
