package helper

import (
	"strconv"
	"unicode"
)

// Stoi parse a string to int, return to default if error.
func Stoi(s string, defaultNum int64) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return defaultNum
	}
	return i
}

// PanicIfErr will panic if param err is not nil
func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

type CharRule struct {
	digit    bool
	letter   bool
	extra    []rune
	min      int
	max      int
	hasUpper bool
	hasToken bool
}

func PasswordCheck(s string, rule CharRule) bool {
	var valid []rune
	if rule.min > 0 && len(s) < rule.min {
		return false
	}
	if rule.max > 0 && len(s) > rule.max {
		return false
	}

	var up bool
	var tok bool

	for _, c := range s {
		if rule.digit && unicode.IsDigit(c) {
			valid = append(valid, c)
		}
		if rule.letter && unicode.IsLetter(c) {
			if unicode.IsUpper(c) {
				up = true
			}
			valid = append(valid, c)
		}
		if rule.extra != nil {
			for _, xc := range rule.extra {
				if xc == c {
					// assume rune in extra is token
					tok = true
					valid = append(valid, c)
					break
				}
			}
		}
	}
	if len(valid) != len(s) {
		return false
	}
	if rule.hasUpper && !up {
		return false
	}
	if rule.hasToken && !tok {
		return false
	}
	return true
}

func AlphaDigitFilter(s string) string {
	var valid []rune

	for _, c := range s {
		if unicode.IsDigit(c) || unicode.IsLetter(c) {
			valid = append(valid, c)
		}
	}
	return string(valid)
}

func AlphaFilter(s string) string {
	var valid []rune

	for _, c := range s {
		if unicode.IsLetter(c) {
			valid = append(valid, c)
		}
	}
	return string(valid)
}

func NumFilter(s string) string {
	var valid []rune

	for _, c := range s {
		if unicode.IsDigit(c) {
			valid = append(valid, c)
		}
	}
	return string(valid)
}
