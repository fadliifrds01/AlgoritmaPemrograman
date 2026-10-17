package main

import (
	"fmt"
	"math"
)

func main() {
	var lantaiMasuk, lantaiKeluar float64

	fmt.Println("Lantai Masuk :")
	fmt.Scan(&lantaiMasuk, &lantaiKeluar)

	jarak := math.Abs(lantaiMasuk - lantaiKeluar)
	fmt.Printf("Jadi hasilnya adalah : %.0f\n", jarak)

}
