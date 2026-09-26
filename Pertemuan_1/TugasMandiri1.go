package main
import "fmt"

func main() {
	var jariJari, totalLuas float64

	fmt.Println("Masukkan Jari Jari : ")
	fmt.Scan(&jariJari)
	phi := 3.14159
	totalLuas = phi * jariJari * jariJari
	fmt.Printf("jadi Luas lingkaran adalah %.02f\n", totalLuas)
}