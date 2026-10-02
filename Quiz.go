package main

import "fmt"

func main() {
	var gram, karat float64

	fmt.Println("Masukkan Berat dalam Gram :")
	fmt.Scan(&gram)

	karat = gram * 5

	fmt.Println("Hasil dalam Karat :", karat)
}