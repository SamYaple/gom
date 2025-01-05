package gom

import (
	"github.com/samyaple/gom/types"
)

func Recognize(parser types.ParserFunc) types.ParserFunc {
	return func(input string) (string, interface{}, error) {
		orig_input := input

		input, _, err := parser(input)
		if err != nil {
			return orig_input, nil, err
		}
		recognized_len := len(orig_input) - len(input)
		return input, orig_input[:recognized_len], nil
	}
}
