package main

import "fmt"

func main() {
	var angka int32

	fmt.Println("Masukkan Angka :")
	fmt.Scan(&angka)

	if angka < 0 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

// Program CekNegatif

// Kamus Data
//     Angka : integer

//     INPUT Angka

//     IF Angka < 0 THEN
//         OUTPUT 1
//     ELSE
//         OUTPUT 0
//     ENDIF

// EndProgram
