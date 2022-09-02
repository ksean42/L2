package pkg

//IsAnagram проверка являются ли строки анаграммами
func IsAnagram(s1 string, s2 string) bool {
	if s1 == "" && s2 == "" {
		return false
	}
	set := make(map[rune]int, len(s1)) // Создаем и заполняем самодельный set c буквой в качестве ключа и количеством букв в качестве значения
	for _, v := range s1 {
		if _, ok := set[v]; !ok {
			set[v] = 1
		} else {
			i := set[v]
			set[v] = i + 1
		}
	}
	// сверяем вторую строку с сетом
	for _, v := range s2 {
		if _, ok := set[v]; ok { // если буква встречалась - отнимаем 1 из значения, иначе - возвращаем false
			i := set[v]
			set[v] = i - 1
		} else {
			return false
		}
	}

	for _, v := range set { // если количество букв в словах отличаются - возвращаем false
		if v != 0 {
			return false
		}
	}

	return true
}
