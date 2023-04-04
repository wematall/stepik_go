package main

import (
	"fmt"
	"strings"
)

func isPalindorm(a, b []string) bool {
	flag := true
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			flag = true
		} else {
			flag = false
			break
		}

	}
	return flag
}

func main() {
	var text string

	fmt.Scan(&text)

	strings.Trim(text, "\n")
	strings.ToLower(text)

	direct := strings.Split(text, "")
	var reversed []string = make([]string, len(direct))
	for i, j := 0, len(direct)-1; i < len(direct); i, j = i+1, j-1 {
		reversed[i] = direct[j]
	}
	fmt.Println(direct)
	fmt.Println(reversed)

	if isPalindorm(direct, reversed) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
