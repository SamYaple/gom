package main

import (
	"fmt"
	"github.com/samyaple/gom"
)

func main() {
	input := "Hello World"
	fmt.Printf("input string: `%s`\n", input)

	parser := gom.Sequence(
		gom.Tag("Hello"),
		gom.Tag(" "),
		gom.Alt(
			gom.Tag("Bob"),
			gom.Tag("World"),
		),
	)
	input, matches, err := parser(input)
	if err != nil {
		fmt.Println("ERROR parsing")
		return
	}

	matchesSlice, ok := matches.([]interface{})
	if !ok {
		fmt.Println("ERROR interface")
		return
	}

	for _, match := range matchesSlice {
		fmt.Printf("matched: `%s`\n", match)
	}
	fmt.Printf("remaining: `%s`\n", input)

	// input string: `Hello World`
	// matched: `Hello`
	// matched: `World`
	// remaining: ``
}
