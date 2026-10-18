
package main

import (
	"fmt"
	"math"
)

func main() {
	var angkaAdik, angkaKakak float64

	fmt.Scan(&angkaAdik, &angkaKakak)

	selisih := math.Abs(angkaAdik - angkaKakak)

	hasil := selisih < 0.0001

	fmt.Println(hasil)
}

// PROGRAM CekKemenanganAdik

// KAMUS DATA
//     angkaAdik, angkaKakak : real
//     selisih : real
//     hasil : boolean

// ALGORITMA
//     input(angkaAdik, angkaKakak)

//     selisih <- ABS(angkaAdik - angkaKakak)

//     hasil <- (selisih < 0.0001)

//     output(hasil)

// ENDPROGRAM