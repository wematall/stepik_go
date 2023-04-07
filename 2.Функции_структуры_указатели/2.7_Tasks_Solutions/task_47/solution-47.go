package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func check_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	var s string
	var s_slice []string
	var err error
	var curr_int int
	max_int := 0

	fmt.Scan(&s)
	s_slice = strings.Split(s, "")

	for i := 0; i < len(s_slice); i++ {
		curr_int, err = strconv.Atoi(s_slice[i])
		check_err(err)
		if max_int <= curr_int {
			max_int = curr_int
		}
	}

	fmt.Println(max_int)
}
