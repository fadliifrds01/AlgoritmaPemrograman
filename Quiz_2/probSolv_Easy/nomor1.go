package main
import "fmt"

func main() {
	var nilai bool

	fmt.Println("Masukkan Nilai :")
	fmt.Scan(&nilai)
	hasil := !nilai
	fmt.Println(hasil)
}

// PROGRAM KebalikanBoolean
// KAMUS DATA
//     nilai, hasil : boolean
// ALGORITMA
//     input(nilai)
//     hasil <- NOT nilai
//     output(hasil)
// ENDPROGRAM