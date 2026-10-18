
package main

import "fmt"

func main() {
	var nilai int

	fmt.Println("Masukkan Nilai :")
	fmt.Scan(&nilai)

	hasil := nilai >= 0 && nilai <= 100
	fmt.Println(hasil)
}

// PROGRAM ValidasiNilai
// KAMUS DATA
//     nilai : integer
//     hasil : boolean

// ALGORITMA
//     input(nilai)
//     hasil <- (nilai >= 0) AND (nilai <= 100)
//     output(hasil)
// ENDPROGRAM