package main

import "fmt"

func main() {
	var tahun int

	fmt.Scan(&tahun)

	hasil := (tahun%400 == 0) || ((tahun%4 == 0) && (tahun%100 != 0))

	fmt.Println(hasil)
}

// PROGRAM TahunKabisat

// KAMUS DATA
//     tahun : integer
//     hasil : boolean

// ALGORITMA
//     input(tahun)

//     hasil <- (tahun mod 400 = 0) OR
//              ((tahun mod 4 = 0) AND (tahun mod 100 != 0))

//     output(hasil)

// ENDPROGRAM