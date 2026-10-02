package main

import "fmt"

func main() {
	var x, y int

	fmt.Scan(&x, &y)

	if x%y == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

// Program CekFaktor

// Kamus Data
//     X : integer
//     Y : integer

//     INPUT X
//     INPUT Y

//     IF X % Y = 0 THEN
//         OUTPUT 1
//     ELSE
//         OUTPUT 0
//     ENDIF

// EndProgram
