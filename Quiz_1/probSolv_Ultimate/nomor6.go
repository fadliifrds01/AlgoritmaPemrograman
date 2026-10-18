package main

import "fmt"

func main() {
	var x, y, z int
	var sementara int

	fmt.Scan(&x, &y, &z)

	sementara = y
	y = x
	x = z
	z = sementara

	fmt.Println(x, y, z)
}

// Program TukarTigaVariabel

// Kamus Data
//     X : integer
//     Y : integer
//     Z : integer
//     Sementara : integer

//     INPUT X
//     INPUT Y
//     INPUT Z

//     Sementara <- Y
//     Y <- X
//     X <- Z
//     Z <- Sementara

//     OUTPUT X
//     OUTPUT Y
//     OUTPUT Z

// EndProgram
