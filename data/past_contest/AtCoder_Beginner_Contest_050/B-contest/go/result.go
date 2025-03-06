package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var t []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputTs := strings.Split(scanner.Text()," ")

	for _, inputT := range inputTs {
		tn, _ := strconv.Atoi(inputT)
		t = append(t, tn)
	}

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	var sumTime int
	for i := 0; i < m; i++ {
		scanner.Scan()
		inputs := strings.Split(scanner.Text(), " ")

		p, _ := strconv.Atoi(inputs[0])
		x, _ := strconv.Atoi(inputs[1])

		sumTime = 0

		for j := 0; j < n; j++ {
			if j == (p - 1) {
				sumTime += x
			} else {
				sumTime += t[j]
			}
		}
		fmt.Println(sumTime)
	}
}
