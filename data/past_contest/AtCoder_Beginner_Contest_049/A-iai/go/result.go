package main

import (
	"fmt"
)

func main() {
	var c string
	check := []string{"a","e","i","o","u"}

	var flag bool = false

	fmt.Scanf("%s", &c)

	for _, v := range check {
		if v == c {
			flag = true
		}
	}

	if flag {
		fmt.Println("vowel")
	} else {
		fmt.Println("consonant")
	}

}