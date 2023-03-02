package word

import (
	"strings"
	"unicode"
)

// ToUpper 转大写格式
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 转小写格式
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Underscore2UpperCamelCase 下划线转大写驼峰
func Underscore2UpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	s = strings.Replace(s, " ", "", -1)

	return s
}

// Underscore2LowerCamelCase 下划线转小写驼峰
func Underscore2LowerCamelCase(s string) string {
	s = Underscore2UpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUnderscore 驼峰转下划线
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
