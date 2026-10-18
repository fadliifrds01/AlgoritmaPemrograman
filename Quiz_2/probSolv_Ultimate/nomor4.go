
package main

import "fmt"

func main() {
	var d1, d2, d3 int

	fmt.Scan(&d1, &d2, &d3)

	hasil := (d1%2 != 0) && (d2%2 != 0) && (d3%2 != 0)

	fmt.Println(hasil)
}

// PROGRAM CekDaduGanjil

// KAMUS DATA
//     d1, d2, d3 : integer
//     hasil : boolean

// ALGORITMA
//     input(d1, d2, d3)

//     hasil <- (d1 mod 2 != 0) AND
//              (d2 mod 2 != 0) AND
//              (d3 mod 2 != 0)

//     output(hasil)

// ENDPROGRAM