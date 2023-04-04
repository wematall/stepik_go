package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var s string
	check := "^[a-zA-Z0-9]{5,}$"

	fmt.Scan(&s)
	strings.Trim(s, "\n")

	match, _ := regexp.MatchString(check, s)

	if match && len(s) >= 5 {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
