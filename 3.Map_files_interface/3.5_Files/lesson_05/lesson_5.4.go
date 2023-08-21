package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	// Все указанные в тексте задачи пакеты уже импортированы (strconv, bufio, os, io)
	// Другие использовать нельзя.
	sum := 0

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		str := string(input)
		num, err := strconv.Atoi(str)
		if err != nil {
			continue
		}

		if num > 100 || num < 0 {
			break
		}

		sum += num
	}

	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(sum))
	writer.Flush()
}
