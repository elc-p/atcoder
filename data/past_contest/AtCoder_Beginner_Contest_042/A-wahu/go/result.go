package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	result := a * b * c

	if result == 175 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}