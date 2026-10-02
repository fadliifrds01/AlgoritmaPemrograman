package main

import "fmt"

func main() {
	var angka1, angka2, angka3, angka4 int
	var x, y int

	fmt.Println("Masukkan Angka :")
	fmt.Scan(&angka1, &angka2, &angka3, &angka4)

	fmt.Println("-- addition --")
	x = angka1*angka4 + angka3*angka2
	y = angka2 * angka4
	fmt.Println("Hasil :", x, "/", y)

	fmt.Println()

	fmt.Println("-- subtraction --")
	x = angka1*angka4 - angka3*angka2
	y = angka2 * angka4
	fmt.Println("Hasil :", x, "/", y)

	fmt.Println()

	fmt.Println("-- multiplication --")
	x = angka1 * angka3
	y = angka2 * angka4
	fmt.Println("Hasil :", x, "/", y)

	fmt.Println()

	fmt.Println("-- division --")
	x = angka1 * angka4
	y = angka2 * angka3
	fmt.Println("Hasil :", x, "/", y)
}

// Program KalkulatorPecahan

// Kamus Data
//     Angka1 : integer
//     Angka2 : integer
//     Angka3 : integer
//     Angka4 : integer
//     X : integer
//     Y : integer

//     INPUT Angka1
//     INPUT Angka2
//     INPUT Angka3
//     INPUT Angka4

//     OUTPUT "-- addition --"
//     X <- Angka1 * Angka4 + Angka3 * Angka2
//     Y <- Angka2 * Angka4
//     OUTPUT X, "/", Y

//     OUTPUT "-- subtraction --"
//     X <- Angka1 * Angka4 - Angka3 * Angka2
//     Y <- Angka2 * Angka4
//     OUTPUT X, "/", Y

//     OUTPUT "-- multiplication --"
//     X <- Angka1 * Angka3
//     Y <- Angka2 * Angka4
//     OUTPUT X, "/", Y

//     OUTPUT "-- division --"
//     X <- Angka1 * Angka4
//     Y <- Angka2 * Angka3
//     OUTPUT X, "/", Y

// EndProgram