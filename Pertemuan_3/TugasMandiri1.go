package main
import "fmt"

type gajiKaryawan struct {
	namaKaryawan string
	gajiPokok int
	presentasePajak int
	pajak int
	gajiBersih int
}

func main() {
	var gaji gajiKaryawan
	
	fmt.Println("Masukkan Nama Karyawan :")
	fmt.Scan(&gaji.namaKaryawan)
	fmt.Println("Masukkan Gaji Pokok :")
	fmt.Scan(&gaji.gajiPokok)
	fmt.Println("Masukkan Presentase Pajak (%) :")
	fmt.Scan(&gaji.presentasePajak)
	gaji.pajak = gaji.gajiPokok * gaji.presentasePajak / 100
	gaji.gajiBersih = gaji.gajiPokok - gaji.pajak

	fmt.Println("")
	fmt.Println("== Gaji Karyawan ==")
	fmt.Println("Nama :", gaji.namaKaryawan)
	fmt.Println("Gaji Pokok :", gaji.gajiPokok)
	fmt.Println("Pajak :", gaji.pajak)
	fmt.Println("Gaji Bersih :", gaji.gajiBersih)
}