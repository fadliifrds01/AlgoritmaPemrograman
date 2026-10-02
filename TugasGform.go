package main
import "fmt"

func satu() {
	fmt.Println("Nomor 1 :")

	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println("Hasil :", a + b)
}

// func dua() {
// 	fmt.Println("Nomor 2")

// 	var a int = 5
// 	var b float64 = 5.0
// 	c := a + b
// 	fmt.Println(c)
// } // Output = Error

func tiga() {
	fmt.Println("Nomor 3 :")

	var a rune = 'A'
	fmt.Println(a)
}

func empat() {
	fmt.Println("Nomor 4 :")

	fmt.Println("5" + "5")
}

func lima() {
	fmt.Println("Nomor 5 :")

	var nama string
	fmt.Println("Masukkan Nama :")
	fmt.Scan(&nama)
	fmt.Println("Hallo", nama)
}

func enam() {
	a := 7 
	b := 2 
	hasil := a / b 
	fmt.Println(hasil) 
}

func tujuh() {
	x := 17 
	y := 5 
	fmt.Println(x % y) 
}

func main() {
	// satu()
	// fmt.Println()
	// dua()
	// fmt.Println()
	// tiga()
	// fmt.Println()
	// empat()
	// fmt.Println()
	// lima()
	// fmt.Println()
	// enam()
	fmt.Println()
	tujuh()
}