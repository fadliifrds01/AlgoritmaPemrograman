package main

import "fmt"

func main() {
	var n int

	fmt.Println("Jumlah Tampil :")
	fmt.Scan(&n)

	fmt.Println("")
	for i := 1; i <= n; i++ {
		ganjil := 2*i - 1
		fmt.Println(ganjil)
	}
}