package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Masukkan Bilangan A :")
	fmt.Scan(&a)
	fmt.Println("Masukkan Bilangan B :")
	fmt.Scan(&b)

	hasil := a%b == 0

	fmt.Println(hasil)
}

// PROGRAM KelipatanBilangan
// KAMUS DATA
//     a, b : integer
//     hasil : boolean

// ALGORITMA
//     input(a, b)
//     hasil <- (a mod b = 0)
//     output(hasil)
// ENDPROGRAM