package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var output strings.Builder
	rs := []rune(input)
	length := utf8.RuneCountInString(input)
	if input == "" {
		return "", nil
	}

	for i := 0; i < length-1; i++ {
		switch {
		case unicode.IsDigit(rs[i]) && i == 0:
			return "", ErrInvalidString

		case unicode.IsDigit(rs[i+1]) && unicode.IsDigit(rs[i]):
			return "", ErrInvalidString

		case unicode.IsDigit(rs[i+1]):
			if num, err := strconv.Atoi(string(rs[i+1])); err == nil {
				buf := strings.Repeat(string(rs[i]), num)
				output.WriteString(buf)
			}

		case !unicode.IsDigit(rs[i]):
			if !unicode.IsDigit(rs[i+1]) {
				output.WriteRune(rs[i])
			}
		}
	}

	if !unicode.IsDigit(rs[length-1]) {
		output.WriteRune(rs[length-1])
	}

	return output.String(), nil
}
