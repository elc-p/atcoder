package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	fmt.Println(n * 800 - (n/15)*200)
}