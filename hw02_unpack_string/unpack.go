package hw02unpackstring

import (
	"errors"
	"strconv"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

type myRune rune

func Unpack(s string) (string, error) {
	ans := make([]rune, 0, len(s))
	runes := []myRune(s)
	if isInt(0, runes) {
		return "", ErrInvalidString
	}
	var curRune myRune
	for i, val := range runes {
		n, ok := val.ConvertToInt()
		if !ok {
			curRune = val
			if isInt(i+1, runes) {
				continue
			}
			ans = append(ans, rune(curRune))
			continue
		}
		if isInt(i+1, runes) {
			return "", ErrInvalidString
		}
		for j := 0; j < n; j++ {
			ans = append(ans, rune(curRune))
		}
	}
	return string(ans), nil
}

// ConvertToInt reports whether the rune is an int digit and returns it.
func (r myRune) ConvertToInt() (int, bool) {
	if unicode.IsDigit(rune(r)) {
		buf := make([]byte, 1)
		_ = utf8.EncodeRune(buf, rune(r))
		v, _ := strconv.Atoi(string(buf))
		return v, true
	}

	return 0, false
}

// isInt returns whether the i-th element of the slice is int.
// If element not found, return false.
func isInt(i int, slice []myRune) bool {
	if i >= len(slice) {
		return false
	}
	_, ok := slice[i].ConvertToInt()
	return ok
}
