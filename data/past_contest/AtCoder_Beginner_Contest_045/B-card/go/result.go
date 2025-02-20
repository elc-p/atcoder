package main

import (
	"fmt"
)

func main() {
	var player int = 0
	var cards []string
	iterator := [3]int{0,0,0}

	var a,b,c string
	fmt.Scanf("%s", &a)
	fmt.Scanf("%s", &b)
	fmt.Scanf("%s", &c)

	cards = append(cards,a)
	cards = append(cards,b)
	cards = append(cards,c)

	for iterator[player] < len(cards[player]) {
		next := cards[player][iterator[player]:iterator[player]+1]
		iterator[player]++
		if next == "a" {
			player = 0
		} else if next == "b" {
			player = 1
		} else {
			player = 2
		}
	}

	if player == 0 {
		fmt.Println("A")
	} else if player == 1 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

}