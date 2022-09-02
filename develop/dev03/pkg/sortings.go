package pkg

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Sort -  Главный объект сортировки
type Sort struct {
	Elems      *[]string
	Operations []Operation
}

// NewSort - конструктор
func NewSort(ops []Operation, elems *[]string) *Sort {
	return &Sort{
		elems,
		ops,
	}
}

// ExecOperations - выполнить все команды из списка
func (s *Sort) ExecOperations() {
	for _, v := range s.Operations {
		v.Exec(s.Elems)
	}
}

// SortByColumn - сортировка по столбцу
type SortByColumn struct {
	Column int
}

//Exec SortByColumn
func (s *SortByColumn) Exec(elems *[]string) {
	sort.Slice(*elems, func(i, j int) bool {
		idx1 := s.Column
		idx2 := s.Column
		str1 := strings.Split((*elems)[i], " ")
		str2 := strings.Split((*elems)[j], " ")
		// Проверки на существование столбца. Если столбца нет - используем последний
		if s.Column > len(str1)-1 {
			idx1 = len(str1) - 1
		}
		res1 := str1[idx1]
		if s.Column > len(str2)-1 {
			idx2 = len(str2) - 1
		}

		res2 := str2[idx2]
		return res1 < res2
	})
}

//SortByNum сортировка по числовому значению
type SortByNum struct {
}

// Exec SortByNum
func (s *SortByNum) Exec(elems *[]string) {
	sort.Slice(*elems, func(i, j int) bool { // Сортируем по числовому значению символа
		reti, _ := strconv.Atoi((*elems)[i])
		retj, _ := strconv.Atoi((*elems)[j])
		return reti < retj
	})
}

//ReverseSort сортировка в обратном порядке
type ReverseSort struct {
}

//Exec ReverseSort
func (s *ReverseSort) Exec(elems *[]string) {
	start := 0
	end := len(*elems) - 1
	for start < end {
		(*elems)[start], (*elems)[end] = (*elems)[end], (*elems)[start]
		start++
		end--
	}
}

//Unique - удалить повторения
type Unique struct {
}

//Exec Unique
func (s *Unique) Exec(elems *[]string) {
	set := make(map[string]struct{}, len(*elems)) // Создаем и наполняем самодельный set
	newSlice := make([]string, 0, len(*elems))
	for _, v := range *elems {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{} // в качестве значений пустая структура размером 0
			newSlice = append(newSlice, v)
		}
	}
	*elems = newSlice
}

//DefaultSort обычная сортировка в лексиграфическом порядке
type DefaultSort struct {
}

//Exec DefaultSort
func (s *DefaultSort) Exec(elems *[]string) {
	sort.Slice(*elems, func(i, j int) bool {
		return (*elems)[i] < (*elems)[j]
	})
}

//IgnoreSpaces игнорируем хвостовые пробелы
type IgnoreSpaces struct {
}

//Exec IgnoreSpaces
func (s *IgnoreSpaces) Exec(elems *[]string) {
	for i := 0; i < len(*elems); i++ {
		(*elems)[i] = strings.TrimRight((*elems)[i], " ") // удаляем хвостовые пробелы
	}
}

// IsSorted - проверка отсортированы ли данные
type IsSorted struct {
}

//Exec IsSorted
func (s *IsSorted) Exec(elems *[]string) {
	for i := 1; i < len(*elems); i++ {
		if (*elems)[i-1] > (*elems)[i] {
			log.Println("disorder:", (*elems)[i])
			break
		}
	}
	os.Exit(0)
}
