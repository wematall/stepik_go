package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	var f *os.File
	file1, _ := os.Create("text.txt")
	file1.WriteString("1")
	file1.WriteString("3")
	switch file1.Name() {
	case "text.txt":
		os.Rename("text.txt", "textxt.txt")
		fallthrough
	case "textxt.txt":
		os.Remove("text.txt")
		fallthrough
	default:
		var err error
		f, err = os.Open("textxt.txt")
		file1.WriteString("2")
		if err != nil {
			panic(err)
		}

	}
	res, _ := ioutil.ReadAll(f)
	resInt, err := strconv.Atoi(string(res[2]))
	if err != nil {
		panic(err)
	}
	fmt.Println(resInt + 1)
	file1.Close()
}
