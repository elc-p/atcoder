package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	
	var ab int = a + b
	var bc int = b + c
	var ca int = c + a

	if (ab == c) || (bc == a) || (ca == b) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
