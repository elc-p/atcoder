package main

import (
	"fmt"
	"math"
)

func calcCost(a []int, ave int, n int) int{
	var cost int = 0
	for i := 0; i < n; i++ {
		cost += int(math.Pow(float64((a[i] - ave)),2))
	}
	return cost
}

func minOutput(ave1 int, ave2 int) int{
	if ave1 >= ave2 {
		return ave2
	} else {
		return ave1
	}
}

func main() {
	var n int
	var a []int
	var total int = 0
	var ave1, ave2 int
	var result int

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var ai int
		fmt.Scanf("%d", &ai)
		a = append(a, ai)
		total += ai
	}

	ave1 = int(math.Ceil(float64(total) / float64(n)))
	ave2 = int(math.Floor(float64(total) / float64(n)))

	result = minOutput(calcCost(a, ave1, n), calcCost(a, ave2, n))

	fmt.Println(result)
}