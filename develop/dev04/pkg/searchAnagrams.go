package pkg

import (
	"sort"
	"strings"
)

func delDuplicates(elems *[]string) {
	set := make(map[string]struct{}, len(*elems))
	newSlice := make([]string, 0, len(*elems)) // создаем новый слайс
	for _, v := range *elems {                 // итерируемся по исходному слайсу и добавляем в новый только уникальные слова
		v = strings.ToLower(v)
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
			newSlice = append(newSlice, v)
		}
	}
	*elems = newSlice //заменяем старый слайс на новый поменяв указатель
}

// SearchAnagrams основная функция поискаа анаграмм
func SearchAnagrams(dict *[]string) map[string][]string {
	delDuplicates(dict)                               // удаляем дубликаты из словаря
	anagrams := make(map[string][]string, len(*dict)) // мапа для хранения анаграм. Ключ - строка с отсортированными буквами , значение - множество анаграм
	result := make(map[string][]string, len(*dict))   // мапа результатов. Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого, слово из множества.

	for _, v := range *dict {
		lowStr := strings.ToLower(v) // приводим к нижнему регистру
		word := []rune(lowStr)       //разбиваем на массив символов и сортируем буквы для удобного сопоставления слов
		sort.Slice(word, func(i, j int) bool {
			return word[i] < word[j]
		})
		anagrams[string(word)] = append(anagrams[string(word)], lowStr)
	}

	for _, v := range anagrams {
		if len(v) == 1 { // пропускаем множества из одного элементы
			continue
		}
		firstWord := v[0] // сохраняем первое встретившееся слово
		//v = v[1:len(v)]
		sort.Slice(v, func(i, j int) bool { // сортируем слова по возрастанию
			return v[i] < v[j]
		})
		result[firstWord] = v // добавляем в мапу
	}

	return result
}
