package main

import (
	"fmt"
)

func numberOfIntegers(a int64, x int64) int64{
	if a <= 0 {
		return 0
	} else {
		return a / x
	}
}

func main() {
	var a, b, x int64
	var toAMinus int64
	var toB int64

	fmt.Scanf("%d %d %d", &a, &b, &x)

	toAMinus = numberOfIntegers(a - 1, x)
	toB = numberOfIntegers(b, x)

	if a == 0 && b >= 0 {
		fmt.Println(toB + 1)
	} else {
		fmt.Println(toB - toAMinus)
	}
	
}