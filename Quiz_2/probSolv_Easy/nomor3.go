package main

import "fmt"

func main() {
	var angka int

	fmt.Println("Masukkan Angka :")
	fmt.Scan(&angka)

	hasil := angka%2 == 0
	fmt.Println(hasil)
}

// PROGRAM BilanganGenap
// KAMUS DATA
//     angka : integer
//     hasil : boolean

// ALGORITMA
//     input(angka)
//     hasil <- (angka mod 2 = 0)
//     output(hasil)
// ENDPROGRAM