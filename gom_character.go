package gom

import (
	"errors"
	"github.com/samyaple/gom/types"
	"unicode"
)

func Alphanumeric1() types.ParserFunc {
	return func(input string) (string, interface{}, error) {
		var idx int
		for idx < len(input) {
			r := rune(input[idx])
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				idx++
			} else {
				break
			}
		}
		if idx == 0 {
			return input, nil, errors.New("alphanumeric1: no match")
		}
		return input[idx:], input[:idx], nil
	}
}
