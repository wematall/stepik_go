package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	result := strings.Trim(text, "\n")
	rs := []rune(result)
	if unicode.IsUpper(rs[0]) && strings.HasSuffix(result, ".") {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
