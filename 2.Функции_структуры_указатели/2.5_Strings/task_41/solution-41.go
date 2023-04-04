package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var old_string, new_string []string

	fmt.Scan(&s)
	strings.Trim(s, "\n")
	old_string = strings.Split(s, "")

	for idx, el := range old_string {
		if idx%2 != 0 {
			new_string = append(new_string, el)
		}
	}
	myString := strings.Join(new_string, "")
	fmt.Println(myString)
}
