package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Masukkan Bilangan A :")
	fmt.Scan(&a)
	fmt.Println("Masukkan Bilangan B :")
	fmt.Scan(&b)

	hasil := a*a == b

	fmt.Println(hasil)
}

// PROGRAM AkarKuadrat
// KAMUS DATA
//     a, b : integer
//     hasil : boolean

// ALGORITMA
//     input(a, b)
//     hasil <- (a * a = b)
//     output(hasil)
// ENDPROGRAM