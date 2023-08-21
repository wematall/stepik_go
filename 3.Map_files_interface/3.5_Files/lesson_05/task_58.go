package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	csvFile, err := os.Open("task.data")
	if err != nil {
		log.Fatal(err)
		fmt.Println("some error")
	}

	defer csvFile.Close()

	rd := bufio.NewReader(csvFile)
	reader := csv.NewReader(rd)
	reader.Comma = ';'

	record, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading records")
	}

	for num, item := range record {
		if item == "0" {
			fmt.Println(num + 1)
		}
	}
}
