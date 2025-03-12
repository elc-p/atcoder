package main

import (
	"fmt"
	"regexp"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)

	pattern := `A.*Z`

	re, _ := regexp.Compile(pattern)

	az := re.FindString(s)

	fmt.Println(len(az))
}