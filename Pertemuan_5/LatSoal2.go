// package main
// import "fmt"

// func main() {
// 	var count int

// 	fmt.Scan(&count)
// 	for i := 1; i <= count; i++ {
// 		for j := 1; j <= count-1; j++ {
// 			fmt.Print(" ")
// 		}

// 		for k := 1; k <= (2*i - 1); k++ {
//             fmt.Print("*")
//         }

// 		fmt.Println()
// 	}
// }

package main
import "fmt"

func main() {
	var count int

	fmt.Scan(&count)
	for i := 1; i <= count; i++ {
		for j := 1; j <= count-i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= (2*i - 1); k++ {
			fmt.Print("*")
		}

		// for l := 4; l <= (2*i - 1); i-- {
		// 	fmt.Println("*")
		// }

		// for m := 1; i < count-1; m++ {
		// 	fmt.Println(" ")
		// }
		
		fmt.Println()

	}

	for i := count; i >= 1 ; i-- {
		for j := 0; j <= count-i; j++ {
			fmt.Print(" ")
		}

		for k := 3; k <= (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
