package main

import "fmt"

func main() {
	var bilangan1, bilangan2 int

	fmt.Println("Masukkan Bilangan 1 :")
	fmt.Scan(&bilangan1)
	fmt.Println("Masukkan Bilangan 2 :")
	fmt.Scan(&bilangan2)

	fmt.Println("")
	for i := bilangan1; i <= bilangan2; i++ {
		fmt.Print(i)
	}
}
