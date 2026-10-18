package main

import "fmt"

func main() {
	var huruf string

	fmt.Scan(&huruf)

	hasil := (huruf >= "a" && huruf <= "z") &&
		(huruf != "a" && huruf != "i" &&
			huruf != "u" && huruf != "e" && huruf != "o")

	fmt.Println(hasil)
}

// PROGRAM CekKonsonan

// KAMUS DATA
//     huruf : string
//     hasil : boolean

// ALGORITMA
//     input(huruf)

//     hasil <- (huruf >= "a" AND huruf <= "z") AND
//              (huruf != "a" AND huruf != "i" AND
//               huruf != "u" AND huruf != "e" AND
//               huruf != "o")

//     output(hasil)

// ENDPROGRAM