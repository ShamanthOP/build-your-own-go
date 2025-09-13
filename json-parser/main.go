package main

import (
	"fmt"
	"jsonparser/jsonparser"
)

func main() {
	input := `
	{
		"a": 123,
		"b": "hello",
		"c": null,
		"d": true
	}
	`
	output, err := jsonparser.Parse(input)

	if err != nil {
		panic(err)
	}

	fmt.Printf("input: %v\nparsed output: %v\n", input, output)
}
