package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var old_string []string
	var new_string []string

	fmt.Scan(&s)
	old_string = strings.Split(s, "")

	for i := 0; i < len(old_string); i++ {
		new_string = append(new_string, old_string[i])

		if i == len(old_string)-1 {
			break
		}
		new_string = append(new_string, "*")
	}

	new_s := strings.Join(new_string, "")

	fmt.Println(new_string)
	fmt.Println(new_s)

}
