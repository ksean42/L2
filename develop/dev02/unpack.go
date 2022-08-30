package unpack

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func unpack(input string) (string, error) {
	str := []rune(input)
	indexes := idx(str)
	var result strings.Builder
	if input == "" {
		return "", nil
	}
	if unicode.IsDigit(str[0]) {
		return "", errors.New("некорректная строка")
	}
	for i, char := range str {
		if v, ok := indexes[i]; ok {
			app := strings.Repeat(string(char), v)
			result.WriteString(app)
			continue
		}
		if char != '\\' && !unicode.IsDigit(char) {
			result.WriteRune(char)
		}
	}
	return result.String(), nil
}

func idx(input []rune) map[int]int {
	var num strings.Builder
	res := make(map[int]int)
	numIdx := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '\\' {
			res[i+1] = 1
			i++
			continue
		}
		if unicode.IsDigit(input[i]) {
			numIdx = i
			for numIdx < len(input) && unicode.IsDigit(input[numIdx]) {
				num.WriteRune(input[numIdx])
				numIdx++
			}
			res[i-1], _ = strconv.Atoi(num.String())
			i = numIdx
			num.Reset()
			continue
		}
	}
	return res
}
