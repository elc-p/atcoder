package main

import (
	"fmt"
)

func check(num int, d []int) bool{
	for 0 < num {
		var mod int = num % 10
		for i := 0; i < len(d); i++{
			if mod == d[i] {
				return false
			}
		}
		num /= 10
	}
	return true
}

func main() {
	var n, k int
	var result int
	fmt.Scanf("%d %d", &n, &k)

	var dk []int
	for i := 0; i < k; i++ {
		var d int
		fmt.Scan(&d)
		dk = append(dk, d)
	}

	for i := n;;i++{
		if check(i,dk) {
			result = i
			break
		}
	}

	fmt.Println(result)
}