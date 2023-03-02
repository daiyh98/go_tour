package word

import (
	"strings"
	"unicode"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func Underscore2UpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	s = strings.Replace(s, " ", "", -1)

	return s
}

func Unerscore2LowerCamelCase(s string) string {
	s = Underscore2UpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}
