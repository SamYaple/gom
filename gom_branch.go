package gom

import (
	"errors"
	"github.com/samyaple/gom/types"
)

func Alt(parsers ...types.ParserFunc) types.ParserFunc {
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
