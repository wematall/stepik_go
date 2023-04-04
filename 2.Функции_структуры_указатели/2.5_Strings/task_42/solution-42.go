package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	counter := 0

	fmt.Scan(&s)
	str := strings.Split(s, "")
	el := str[0]

	myString := make([]string, 0)

	for i := 0; i < len(str); i++ {
		el = str[i]
		counter = 0
		for j := 0; j < len(str); j++ {
			if el == str[j] {
				counter += 1
			}
		}

		if counter == 1 {
			myString = append(myString, el)
		}

	}

	res := strings.Join(myString, "")

	fmt.Println(res)

}
