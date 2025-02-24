package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k int
	var pattern int
	fmt.Scanf("%d %d", &n, &k)

	if n == 1 {
		pattern = k
	} else {
		pattern = k * int(math.Pow(float64(k - 1), float64(n - 1)))
	}

	fmt.Println(pattern)
}
