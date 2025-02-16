package main

import (
	"fmt"
)

func main() {
	var n, k, x, y int
	var payment int = 0

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &k)
	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &y)

	if n - k > 0 {
		payment = k * x + (n - k) * y
	} else {
		payment = n * x
	}
	

	fmt.Println(payment)
}