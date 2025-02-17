package main

import (
	"fmt"
	"strings"
)

func main() {
	var w string
	var alphabets = [26]string {"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	var flag int = 1

	fmt.Scanf("%s", &w)

	for i := 0; i < 26; i++ {
		if strings.Count(w,alphabets[i]) % 2 != 0 {
			flag = 0
			break
		}
	}

	if flag == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}