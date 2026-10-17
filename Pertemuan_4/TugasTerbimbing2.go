package main
import "fmt"

func main() {
	var bilPertama int
	total := 1
	fmt.Println("Masukkan Bilangan Bulat:")
	fmt.Scan(&bilPertama)
	for i := 1; i <= bilPertama; i++ {
		total *= i
	}
	fmt.Print("Hasil : ", total)
}	