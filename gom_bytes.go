package gom

import (
	"errors"
	"github.com/samyaple/gom/types"
	"strings"
)

func Tag(tag string) types.ParserFunc {
	return func(input string) (string, interface{}, error) {
		if strings.HasPrefix(input, tag) {
			return input[len(tag):], tag, nil
		}
		return input, nil, errors.New("tag does not match")
	}
}
