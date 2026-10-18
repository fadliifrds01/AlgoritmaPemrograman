package main

import "fmt"

func main() {
	var angka int

	fmt.Println("Masukkan Angka :")
	fmt.Scan(&angka)

	hasil := angka < 0
	fmt.Println(hasil)
}

// PROGRAM BilanganNegatif
// KAMUS DATA
//     angka : integer
//     hasil : boolean

// ALGORITMA
//     input(angka)
//     hasil <- (angka < 0)
//     output(hasil)
// ENDPROGRAM