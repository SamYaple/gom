package gom

import (
	"github.com/samyaple/gom/types"
)

func Tuple(parsers ...types.ParserFunc) types.ParserFunc {
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
