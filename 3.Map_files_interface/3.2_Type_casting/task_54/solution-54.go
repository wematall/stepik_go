package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var s_new, s string
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}

	// s = strings.Trim(line, "\n")
	s = strings.ReplaceAll(line, " ", "")
	s_new = strings.ReplaceAll(s, ",", ".")
	s_new = strings.ReplaceAll(s_new, ";", " ")
	s_new = strings.Trim(s_new, "\n")
	str_arr := strings.Split(s_new, " ")

	a, err := strconv.ParseFloat(str_arr[0], 64)
	if err != nil {
		panic(err)
	}
	b, err := strconv.ParseFloat(str_arr[1], 64)
	if err != nil {
		panic(err)
	}
	res := a / b

	fmt.Printf("%.4f \n", res)
}
