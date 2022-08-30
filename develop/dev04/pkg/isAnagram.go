package pkg

func IsAnagram(s1 string, s2 string) bool {
	if s1 == "" && s2 == "" {
		return false
	}
	set := make(map[rune]int, len(s1))
	for _, v := range s1 {
		if _, ok := set[v]; !ok {
			set[v] = 1
		} else {
			i := set[v]
			set[v] = i + 1
		}
	}

	for _, v := range s2 {
		if _, ok := set[v]; ok {
			i := set[v]
			set[v] = i - 1
		} else {
			return false
		}
	}

	for _, v := range set {
		if v != 0 {
			return false
		}
	}

	return true
}
