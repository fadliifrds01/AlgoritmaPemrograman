package main

import "fmt"

func main() {
	var karakter string

	fmt.Scan(&karakter)

	hasil := (karakter == "a") || (karakter == "i") ||
		(karakter == "u") || (karakter == "e") ||
		(karakter == "o")

	fmt.Println(hasil)
}

// PROGRAM CekVokal

// KAMUS DATA
//     karakter : string
//     hasil : boolean

// ALGORITMA
//     input(karakter)

//     hasil <- (karakter = "a") OR
//              (karakter = "i") OR
//              (karakter = "u") OR
//              (karakter = "e") OR
//              (karakter = "o")

//     output(hasil)

// ENDPROGRAM