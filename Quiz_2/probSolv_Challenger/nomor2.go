
package main

import "fmt"

func main() {
	var panjang, lebar, tinggi, berat int

	fmt.Scan(&panjang, &lebar, &tinggi, &berat)

	volume := panjang * lebar * tinggi
	hasil := (volume <= 1000000) && (berat <= 30000)

	fmt.Println(hasil)
}

// PROGRAM PengirimanPaket

// KAMUS DATA
//     panjang, lebar, tinggi, berat : integer
//     volume : integer
//     hasil : boolean

// ALGORITMA
//     input(panjang, lebar, tinggi, berat)

//     volume <- panjang * lebar * tinggi

//     hasil <- (volume <= 1000000) AND (berat <= 30000)

//     output(hasil)

// ENDPROGRAM