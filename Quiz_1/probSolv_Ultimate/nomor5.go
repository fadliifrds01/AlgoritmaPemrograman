package main

import "fmt"

func main() {
	var totalDetik int

	fmt.Println("Masukkan waktu dalam detik :")
	fmt.Scan(&totalDetik)

	jam := totalDetik / 3600
	sisa := totalDetik % 3600

	menit := sisa / 60
	detik := sisa % 60

	fmt.Println("Jam   :", jam)
	fmt.Println("Menit :", menit)
	fmt.Println("Detik :", detik)
}

// Program KonversiWaktu

// Kamus Data
//     TotalDetik : integer
//     Jam : integer
//     Sisa : integer
//     Menit : integer
//     Detik : integer

//     INPUT TotalDetik

//     Jam <- TotalDetik / 3600
//     Sisa <- TotalDetik % 3600

//     Menit <- Sisa / 60
//     Detik <- Sisa % 60

//     OUTPUT Jam
//     OUTPUT Menit
//     OUTPUT Detik

// EndProgram
