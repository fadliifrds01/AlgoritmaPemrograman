
package main

import (
	"fmt"
	"math"
)

func main() {
	var angkaAdik, angkaKakak int

	fmt.Scan(&angkaAdik, &angkaKakak)

	selisih := int(math.Abs(float64(angkaAdik - angkaKakak)))

	hasil := (selisih == 2) || (selisih == 3) ||
		(selisih == 5) || (selisih == 7)

	fmt.Println(hasil)
}

// PROGRAM CekBilanganPrima

// KAMUS DATA
//     angkaAdik, angkaKakak : integer
//     selisih : integer
//     hasil : boolean

// ALGORITMA
//     input(angkaAdik, angkaKakak)

//     selisih <- ABS(angkaAdik - angkaKakak)

//     hasil <- (selisih = 2) OR
//              (selisih = 3) OR
//              (selisih = 5) OR
//              (selisih = 7)

//     output(hasil)

// ENDPROGRAM