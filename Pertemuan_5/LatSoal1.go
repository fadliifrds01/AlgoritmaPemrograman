package main

import "fmt"

func main() {
	var count int
	fmt.Scan(&count)
	for i := count; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}