package main

import "fmt"

func main() {
	var kecepatan int

	fmt.Println("Masukkan Kecepatan :")
	fmt.Scan(&kecepatan)

	totalJarak := 330

	totalWaktu := (totalJarak * 60) / kecepatan

	jam := totalWaktu / 60
	menit := totalWaktu % 60

	fmt.Println(jam, "Jam", menit, "Menit")
}
