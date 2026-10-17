package main

import "fmt"

func main() {
	var bilangan1, bilangan2 int
	var hasil int = 0

	fmt.Println("Masukkan Bilangan 1 :")
	fmt.Scan(&bilangan1)
	fmt.Println("Masukkan Bilangan 2 :")
	fmt.Scan(&bilangan2)

	for i := 1; i <= bilangan2; i++ {
		hasil = hasil + bilangan1
	}
	fmt.Println(hasil)
}
