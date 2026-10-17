package main
import "fmt"

func main() {
	var farenhit, celcius int

	fmt.Println("Masukkan suhu celcius yang di konversi :")
	fmt.Scan(&celcius)
	farenhit = (celcius * 9/5) + 32
	fmt.Println("Jadi suhu celcius yang sudah di konversi adalah :", farenhit, "°F")
}