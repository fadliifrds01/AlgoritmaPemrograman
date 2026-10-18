
package main

import "fmt"

func main() {
	var drachma int

	fmt.Scan(&drachma)

	totalOboli := drachma * 6

	atticTalent := totalOboli / 36000
	sisaOboli := totalOboli % 36000

	mine := sisaOboli / 600
	oboli := sisaOboli % 600

	fmt.Println(atticTalent, mine, oboli)
}

// PROGRAM KonversiDrachma

// KAMUS DATA
//     drachma, totalOboli : integer
//     atticTalent, mine, oboli : integer
//     sisaOboli : integer

// ALGORITMA
//     input(drachma)

//     totalOboli <- drachma * 6

//     atticTalent <- totalOboli div 36000
//     sisaOboli <- totalOboli mod 36000

//     mine <- sisaOboli div 600
//     oboli <- sisaOboli mod 600

//     output(atticTalent, mine, oboli)

// ENDPROGRAM