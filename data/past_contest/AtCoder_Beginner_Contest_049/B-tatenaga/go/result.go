package main

import (
	"fmt"
)

func main() {
	var h, w int
	fmt.Scanf("%d %d", &h, &w)

	// c := make([][]string, h, w)
	var c string

	for i := 0; i < h; i++ {
		fmt.Scanf("%s", &c)
		fmt.Println(c)
		fmt.Println(c)
	}
}