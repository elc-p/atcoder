package main

import (
	"fmt"
)

func main() {
	var a, b int
	var op string

	fmt.Scanf("%d %s %d", &a, &op, &b)

	if op == "+" {
		fmt.Println(a + b)
	} else {
		fmt.Println(a - b)
	}
}
