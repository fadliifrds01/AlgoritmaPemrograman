package main
import "fmt"

func main() {
	var nasiGoreng, ayamGeprek, hargaNasiGoreng, hargaAyamGeprek int

	fmt.Println("Masukkan harga nasi goreng :")
	fmt.Scan(&hargaNasiGoreng)
	fmt.Println("Masukkan harga Ayam Geprek :")
	fmt.Scan(&hargaAyamGeprek)
	fmt.Println("Masukkan Jumlah nasi goreng")
	fmt.Scan(&nasiGoreng)
	fmt.Println("Masukkan Jumlah Ayam Geprek")
	fmt.Scan(&ayamGeprek)

	total := (hargaNasiGoreng * nasiGoreng) + (hargaAyamGeprek * ayamGeprek)
	fmt.Println("Total Biaya Makan Adalah :", total)
}