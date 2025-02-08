package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var n, l int
	fmt.Scanf("%d %d", &n, &l)
	
	var s []string
	for i := 0; i < n; i++ {
		var st string
		fmt.Scan(&st)
		s = append(s,st)
	}
	sort.Strings(s)
	fmt.Println(strings.Join(s,""))
}
