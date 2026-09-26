package main
import "fmt"

func main(){
	var biayaParkir, jam int

	fmt.Println("Masukkan Durasi Jam Parkir :")
	fmt.Scan(&jam)
	biayaParkir = 3000 + (1000 * jam)
	fmt.Println("Biaya Akhir Parkir Adalah :", biayaParkir)
}