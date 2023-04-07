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
	var a, tmp int
	var s, s_tmp string
	var s_slice []string   // input slice
	var out_slice []string // output slice
	var err error

	fmt.Scan(&a)

	s = strconv.Itoa(a)
	s_slice = strings.Split(s, "")

	for i := 0; i < len(s_slice); i++ {
		tmp, err = strconv.Atoi(s_slice[i])
		check_err(err)
		tmp = tmp * tmp
		s_tmp = strconv.Itoa(tmp)

		out_slice = append(out_slice, s_tmp)
	}

	s = strings.Join(out_slice, "")
	a, err = strconv.Atoi(s)
	check_err(err)

	fmt.Println(a)
}
