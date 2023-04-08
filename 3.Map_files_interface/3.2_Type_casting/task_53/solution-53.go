package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func adding(string1, string2 string) int64 {
	a1 := []rune(string1)
	a2 := []rune(string2)
	var s1, s2 string

	for i := 0; i < len(a1); i++ {
		if unicode.IsDigit(a1[i]) {
			s1 += string(a1[i])
		}
	}

	for _, el := range a2 {
		if unicode.IsDigit(el) {
			s2 += string(el)
		}
	}
	a, _ := strconv.ParseInt(s1, 10, 64)
	b, _ := strconv.ParseInt(s2, 10, 64)

	return a + b

}

func main() {
	var a, b string
	fmt.Scan(&a)
	fmt.Scan(&b)

	res := adding(a, b)
	fmt.Println(res)
}
