package pkg

import (
	"sort"
	"strings"
)

func delDuplicates(elems *[]string) {
	set := make(map[string]struct{}, len(*elems))
	newSlice := make([]string, 0, len(*elems))
	for _, v := range *elems {
		v = strings.ToLower(v)
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
			newSlice = append(newSlice, v)
		}
	}
	*elems = newSlice
}

func SearchAnagrams(dict *[]string) map[string][]string {
	delDuplicates(dict)
	anagrams := make(map[string][]string, len(*dict))
	result := make(map[string][]string, len(*dict))
	for _, v := range *dict {
		lowStr := strings.ToLower(v)
		word := []rune(lowStr)
		sort.Slice(word, func(i, j int) bool {
			return word[i] < word[j]
		})
		anagrams[string(word)] = append(anagrams[string(word)], lowStr)
	}

	for _, v := range anagrams {
		if len(v) == 1 {
			continue
		}
		firstWord := v[0]
		//v = v[1:len(v)]
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
		result[firstWord] = v
	}

	return result
}
