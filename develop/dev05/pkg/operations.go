package pkg

import (
	"fmt"
	"regexp"
)

// Operation Интерфейс стратегии
type Operation interface {
	Exec(elems []string, pattern string)
}

// Count Конкретная стратегия подсчета строк
type Count struct {
}

// Exec Функция подсчета строк удовлетворяющих паттерну
func (f *Count) Exec(elems []string, pattern string) {
	count := 0

	for _, v := range elems {
		y, err := regexp.MatchString(pattern, v)
		if err == nil && y {
			count++
		}
	}
	fmt.Println(count)
}

// Context Конкретная стратегия печати строк до и после совпадения
type Context struct {
	Before int
	After  int
}

// Exec Функция печати строк до и после строки удовлетворяющей паттерну
func (f *Context) Exec(elems []string, pattern string) {
	for i := 0; i < len(elems); i++ {
		y, err := regexp.MatchString(pattern, elems[i])
		if err == nil && y {
			var start int
			if i-f.Before < 0 {
				start = 0
			} else {

				start = i - f.Before
			}
			for start < len(elems) && start < i+f.After+1 {
				fmt.Println(elems[start])
				start++
			}
			fmt.Println("___________")
		}
	}
}

// Invert Конкретная стратегия печати строк не удовлетворяющих паттерну
type Invert struct {
}

// Exec Функция печати строк не удовлетворяющих паттерну
func (f *Invert) Exec(elems []string, pattern string) {
	for i := 0; i < len(elems); i++ {
		y, err := regexp.MatchString(pattern, elems[i])
		if err == nil && !y {
			fmt.Println(elems[i])
		}
	}
}

// Fixed Конкретная стратегия печати строк точно совпадающик с паттерном
type Fixed struct {
}

// Exec Функция печати строк точно совпадающик с паттерном
func (f *Fixed) Exec(elems []string, pattern string) {
	for i := 0; i < len(elems); i++ {
		if elems[i] == pattern {
			fmt.Println(elems[i])
			fmt.Println("___________")
		}
	}
}

// NumOfLine Конкретная стратегия печати номеров строк удовлетворяющих паттерну
type NumOfLine struct {
}

// Exec Функция печати номеров строк удовлетворяющих паттерну
func (f *NumOfLine) Exec(elems []string, pattern string) {
	for i := 0; i < len(elems); i++ {
		y, err := regexp.MatchString(pattern, elems[i])
		if err == nil && y {
			fmt.Println(i + 1)
		}
	}
}
