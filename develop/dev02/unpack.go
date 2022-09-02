package unpack

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func unpack(input string) (string, error) {
	str := []rune(input)
	indexes := idx(str) // создаем мапу где в качестве ключа - индекс символа, а в качестве значения - число повторений
	var result strings.Builder

	if input == "" { // Если подана пустая строка - возвращаем пустую строку
		return "", nil
	}
	if unicode.IsDigit(str[0]) { //ошибка в случае некорректной строки
		return "", errors.New("некорректная строка")
	}
	for i, char := range str {
		if v, ok := indexes[i]; ok { // проходим по массиву и сверяемся с мапой
			app := strings.Repeat(string(char), v) // если индекс есть в мапе - повторяем нужное кол-во раз и записываем в итоговую строку
			result.WriteString(app)
			continue
		}
		if char != '\\' && !unicode.IsDigit(char) { // записываем цифру если она экранирована
			result.WriteRune(char)
		}
	}
	return result.String(), nil
}

// Возвращает мапу: в качестве ключа - индекс символа, а в качестве значения - число повторений
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
