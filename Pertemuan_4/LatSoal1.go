package main
import "fmt"

func main() {
	var hp, attack int

	fmt.Println("Masukkan HP :")
	fmt.Scan(&hp)
	fmt.Println("Masukkan Attack :")
	fmt.Scan(&attack)

	fmt.Println("")
	for i := hp; i >= 0; i-=attack {
		fmt.Println("HP Karakter Saat Ini :", i)
	}

}