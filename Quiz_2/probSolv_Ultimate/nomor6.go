
package main

import (
	"fmt"
	"math"
)

func main() {
	var angkaAdik, angkaKakak int

	fmt.Scan(&angkaAdik, &angkaKakak)

	selisih := math.Abs(float64(angkaAdik - angkaKakak))

	hasil := (angkaAdik == angkaKakak) ||
		(selisih == 1) ||
		(selisih == 5)

	fmt.Println(hasil)
}

// PROGRAM CekKemenangan

// KAMUS DATA
//     angkaAdik, angkaKakak : integer
//     selisih : real
//     hasil : boolean

// ALGORITMA
//     input(angkaAdik, angkaKakak)

//     selisih <- ABS(angkaAdik - angkaKakak)

//     hasil <- (angkaAdik = angkaKakak) OR
//              (selisih = 1) OR
//              (selisih = 5)

//     output(hasil)

// ENDPROGRAM