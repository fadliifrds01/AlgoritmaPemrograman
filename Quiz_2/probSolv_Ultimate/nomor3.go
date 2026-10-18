
package main

import (
	"fmt"
	"math"
)

func main() {
	var d1, d2, d3 int

	fmt.Scan(&d1, &d2, &d3)

	hasil := (math.Abs(float64(d1-d2)) == 1) &&
		(math.Abs(float64(d2-d3)) == 1)

	fmt.Println(hasil)
}

// PROGRAM DaduBerurutan

// KAMUS DATA
//     d1, d2, d3 : integer
//     hasil : boolean

// ALGORITMA
//     input(d1, d2, d3)

//     hasil <- (ABS(d1 - d2) = 1) AND
//              (ABS(d2 - d3) = 1)

//     output(hasil)

// ENDPROGRAM