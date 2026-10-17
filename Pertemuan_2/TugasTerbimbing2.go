
package main

import "fmt"

func main() {
	var kedalaman, batas float64
	const masaJenis float64 = 1025
	const percepatan float64 = 9.8

	fmt.Println("Masukkan Kedalaman Menyelam :")
	fmt.Scan(&kedalaman)
	fmt.Println("Masukkan Batas :")
	fmt.Scan(&batas)

	tekanan := masaJenis * percepatan * kedalaman
	aman := tekanan <= batas
	fmt.Printf("%.1f\n", tekanan)
	fmt.Println(aman)
}