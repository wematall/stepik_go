package main

// что выведет программа после запуска?
import (
	"encoding/json"
	"fmt"
)

func main() {
	data1 := []byte(`
	{"menu": {
		"id": "file",
		"value": "File",
		"popup": {
			"menuitem": [
				{"value": "New", "onclick": "CreateNewDoc()"},
				{"value": "!", "onclick": "Doc()"},
				{"value": "Close", "onclick": "CloseDoc()"}
			]
		}
	}}`)

	data2 := []byte(`
		{"text": {
			"id": "file",
			"value": "kek",
			"popup": {
				"menuitem": [
					{"value": "new2", "power": "CreateNew()"},
					{"value": "new" "onclick": "Close()"}
				]
			}
		}}`)

	fmt.Print(json.Valid(data1))
	fmt.Print(" ")
	fmt.Print(json.Valid(data2))
}
