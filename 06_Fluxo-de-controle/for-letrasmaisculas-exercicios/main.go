package main

import (
	"fmt"
)

// - Desafio surpresa!
// - Format printing:
//     - Decimal       %d
//     - Hexadecimal   %#x
//     - Unicode       %#U
//     - Tab           \t
//     - Linha nova    \n
// - Faça um loop dos números 33 a 122, e utilize format printing para demonstrá-los como texto/string.

func main() {
	
	
	for x := 65; x <= 90; x++ {
		fmt.Printf("%v:\n", x)
		for i := 0; i < 3; i++ {
			fmt.Printf("\t%#U\n", x)
		}
	}
}
