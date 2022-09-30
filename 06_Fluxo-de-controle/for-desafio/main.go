package main

import (
	"fmt"
)


func main() {
	x := 10
	
	for {
		if x >= 20 {
			break
		}
		fmt.Println(x)
		x++
	}
}
