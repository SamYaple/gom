package gom

import (
	"errors"
	"strings"
)

type ParserFunc func(input string) (string, interface{}, error)

func Tag(tag string) ParserFunc {
	return func(input string) (string, interface{}, error) {
		if strings.HasPrefix(input, tag) {
			return input[len(tag):], tag, nil
		}
		return input, nil, errors.New("tag does not match")
	}
}

func Sequence(parsers ...ParserFunc) ParserFunc {
	return func(input string) (string, interface{}, error) {
		orig_input := input
		results := make([]interface{}, 0, len(parsers))
		for _, parser := range parsers {
			remaining, result, err := parser(input)
			if err != nil {
				return orig_input, nil, err
			}
			results = append(results, result)
			input = remaining
		}
		return input, results, nil
	}
}

func Alt(parsers ...ParserFunc) ParserFunc {
	return func(input string) (string, interface{}, error) {
		for _, parser := range parsers {
			input, result, err := parser(input)
			if err != nil {
				continue
			}
			return input, result, nil
		}
		return input, nil, errors.New("alt parsers do not match")
	}
}

//
//func main() {
//    input := "Hello World"
//    fmt.Printf("input string: `%s`\n", input)
//
//    parser := Sequence(
//        Tag("Hello"),
//	Alt(
//	    Tag(" Bob"),
//	    Tag(" World"),
//	),
//    )
//    input, matches, err := parser(input)
//    if err != nil {
//        fmt.Println("ERROR parsing")
//	return
//    }
//
//    matchesSlice, ok := matches.([]interface{})
//    if !ok {
//        fmt.Println("ERROR interface")
//	return
//    }
//
//    for _, match := range matchesSlice {
//        fmt.Printf("matched: `%s`\n", match)
//    }
//    fmt.Printf("remaining: `%s`\n", input)
//}
