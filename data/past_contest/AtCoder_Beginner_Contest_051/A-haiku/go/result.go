package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)

	fmt.Println(strings.Replace(s, ",", " ", -1))
}