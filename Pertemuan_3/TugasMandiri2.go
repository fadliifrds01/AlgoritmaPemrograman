package main
import "fmt"

type mahasiswa struct {
	uts float64
	uas float64
	tugas float64
}

func main() {
	var mhs mahasiswa

	fmt.Println("Masukkan Nilai UTS :")
	fmt.Scan(&mhs.uts)
	fmt.Println("Masukkan Nilai UAS :")
	fmt.Scan(&mhs.uas)
	fmt.Println("Masukkan Nilai Tugas :")
	fmt.Scan(&mhs.tugas)

	nilaiAkhir := mhs.uts*0.30 + mhs.uas*0.40 + mhs.tugas*0.30
	fmt.Println("Niali Akhirnya :", nilaiAkhir)
}