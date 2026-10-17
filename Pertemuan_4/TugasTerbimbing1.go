package main
import "fmt"

func main() {
	var angka, pangkat int
	var hasil int = 1

	fmt.Println("Masukkan Angka :")
	fmt.Scan(&angka)
	fmt.Println("Masukkan Pangkat :")
	fmt.Scan(&pangkat)

	for i := 1; i <= pangkat; i++ {
		hasil = hasil * angka
	}
	fmt.Println(hasil)
}
