package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	var unique = []int{}

	fmt.Scanf("%d %d %d", &a, &b, &c)

	unique = append(unique, a)

	if a != b {
		unique = append(unique, b)
	}

	if a != c && b != c {
		unique = append(unique, c)
	}

	fmt.Println(len(unique))
}