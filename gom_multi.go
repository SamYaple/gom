package gom

import (
	"errors"
	"github.com/samyaple/gom/types"
)

func Many1(parser types.ParserFunc) types.ParserFunc {
	return func(input string) (string, interface{}, error) {
		var results []interface{}
		orig_input := input

		for {
			remaining, result, err := parser(input)
			if err != nil {
				if len(results) == 0 {
					return orig_input, nil, err
				}
				break
			}
			if remaining == input {
				results = append(results, result)
				break
			}
			results = append(results, result)
			input = remaining
		}
		if len(results) == 0 {
			return orig_input, nil, errors.New("many1: zero matches")
		}

		return input, results, nil
	}
}
