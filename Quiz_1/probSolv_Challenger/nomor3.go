package main

import "fmt"

func main() {
	var x, f float64

	fmt.Println("Masukkan nilai x :")
	fmt.Scan(&x)

	f = (x*x*x + 3*x) / (x*x*x*x - 3*x*x + 4)

	fmt.Printf("Hasil f(x) : %.6f\n", f)
}

// Program FungsiFx

// Kamus Data
//     X : real
//     F : real

//     INPUT X
//     F <- (X * X * X + 3 * X) / (X * X * X * X - 3 * X * X + 4)
//     OUTPUT F

// EndProgram