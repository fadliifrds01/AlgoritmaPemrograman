
package main

import "fmt"

func main() {
	var health, score, level int

	fmt.Scan(&health, &score, &level)

	lanjut := (health > 0) &&
		(((level == 1) && (score >= 1000)) ||
			((level == 2) && (score >= 3000)) ||
			(score >= 7000))

	gameOver := !lanjut

	fmt.Println(gameOver)
}

// PROGRAM FallGuysThree

// KAMUS DATA
//     health, score, level : integer
//     lanjut, gameOver : boolean

// ALGORITMA
//     input(health, score, level)

//     lanjut <- (health > 0) AND
//               (((level = 1) AND (score >= 1000)) OR
//                ((level = 2) AND (score >= 3000)) OR
//                (score >= 7000))

//     gameOver <- NOT lanjut

//     output(gameOver)

// ENDPROGRAM