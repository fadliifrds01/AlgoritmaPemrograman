package main
import "fmt"

type Orang struct {
	nama string
	umur int
	tglLahir int
	blnLahir int
	thnLahir int
}

func main() {
	var org Orang

	fmt.Println("Masukkan Nama :")
	fmt.Scan(&org.nama)
	fmt.Println("Masukkan Umur :")
	fmt.Scan(&org.umur)
	fmt.Println("Masukkan Tanggal Lahir :")
	fmt.Scan(&org.tglLahir)
	fmt.Println("Masukkan Bulan Lahir :")
	fmt.Scan(&org.blnLahir)
	fmt.Println("Masukkan Tahun Lahir :")
	fmt.Scan(&org.thnLahir)
}