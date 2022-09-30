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
	
	
	for x := 33; x <= 122; x++ {

		fmt.Printf("%d\t", x)
		fmt.Printf("%#x\t", x)
		fmt.Printf("%v\t", string(x))
		fmt.Printf("%#U\n", x)
	}
}
