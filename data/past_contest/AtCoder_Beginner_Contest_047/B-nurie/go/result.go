package main

import (
	"fmt"
)

func min(x int, y int) int{
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x int, y int) int{
	if x < y {
		return y
	} else {
		return x
	}
}

func main() {
	var w, h, n int

	fmt.Scanf("%d %d %d", &w, &h, &n)

	var xsup, xinf, ysup, yinf int = w, 0, h, 0

	coordinates := make([][]int, n)

	for i := range coordinates {
		coordinates[i] = make([]int, 3)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scanf("%d", &coordinates[i][j])
		}
		if coordinates[i][2] == 1 {
			xinf = max(coordinates[i][0],xinf)
		} else if coordinates[i][2] == 2 {
			xsup = min(coordinates[i][0],xsup)
		} else if coordinates[i][2] == 3 {
			yinf = max(coordinates[i][1],yinf)
		} else {
			ysup = min(coordinates[i][1],ysup)
		}
	}

	if xinf < xsup && yinf < ysup {
		fmt.Println( (xsup - xinf)*(ysup - yinf) )
	} else {
		fmt.Println( 0 )
	}

}
