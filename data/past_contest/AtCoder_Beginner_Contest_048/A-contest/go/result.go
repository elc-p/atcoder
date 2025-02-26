package main

import (
	"fmt"
)

func main() {
	var a, s, c string
	fmt.Scanf("%s %s %s", &a, &s, &c)

	fmt.Println("A" + s[0:1] + "C")
}
