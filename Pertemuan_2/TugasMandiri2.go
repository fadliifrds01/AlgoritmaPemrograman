package main

import "fmt"

func main() {
	var jam, menit, detik int

	fmt.Println("Masukkan Jam :")
	fmt.Scan(&jam)
	fmt.Println("Masukkan Menit :")
	fmt.Scan(&menit)
	fmt.Println("Masukkan Detik :")
	fmt.Scan(&detik)

	totalDetik := (jam * 3600) + (menit * 60) + detik
	fmt.Println("Jadi total detiknya adalah :", totalDetik)
}
