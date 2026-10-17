package main

import (
	"fmt"
	"math"
)

func main() {
	var jariJari, keliling, luas float64

	fmt.Println("Masukkan Jari-Jari :")
	fmt.Scan(&jariJari)

	phi := 3.14
	keliling = 2 * phi * jariJari
	luas = phi * math.Pow(jariJari, 2)

	fmt.Printf("Keliling: %.2f, Luas: %.2f\n", keliling, luas)
}
