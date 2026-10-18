package main

import "fmt"

func main() {
	var x, y float64

	fmt.Println("Masukkan nilai x :")
	fmt.Scan(&x)

	y = (x*x + 1/x) * (x*x + 1/x)

	fmt.Printf("Hasil y : %.6f\n", y)
}

// Program HitungY

// Kamus Data
//     X : real
//     Y : real

//     INPUT X

//     Y <- (X * X + 1 / X) * (X * X + 1 / X)

//     OUTPUT Y

// EndProgram