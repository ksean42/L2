package pkg

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Главный объект сортировки
type Sort struct {
	Elems      *[]string
	Operations []Operation
}

func NewSort(ops []Operation, elems *[]string) *Sort {
	return &Sort{
		elems,
		ops,
	}
}

func (s *Sort) ExecOperations() {
	for _, v := range s.Operations {
		v.Exec(s.Elems)
	}
}

type SortByColumn struct {
	Column int
}

func (s *SortByColumn) Exec(elems *[]string) {
	sort.Slice(*elems, func(i, j int) bool {
		idx1 := s.Column
		idx2 := s.Column
		str1 := strings.Split((*elems)[i], " ")
		str2 := strings.Split((*elems)[j], " ")
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

type SortByNum struct {
}

func (s *SortByNum) Exec(elems *[]string) {
	sort.Slice(*elems, func(i, j int) bool {
		reti, _ := strconv.Atoi((*elems)[i])
		retj, _ := strconv.Atoi((*elems)[j])
		return reti < retj
	})
}

type ReverseSort struct {
}

func (s *ReverseSort) Exec(elems *[]string) {
	start := 0
	end := len(*elems) - 1
	for start < end {
		(*elems)[start], (*elems)[end] = (*elems)[end], (*elems)[start]
		start++
		end--
	}
}

type Unique struct {
}

func (s *Unique) Exec(elems *[]string) {
	set := make(map[string]struct{}, len(*elems))
	newSlice := make([]string, 0, len(*elems))
	for _, v := range *elems {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
			newSlice = append(newSlice, v)
		}
	}
	*elems = newSlice
}

type DefaultSort struct {
}

func (s *DefaultSort) Exec(elems *[]string) {
	sort.Slice(*elems, func(i, j int) bool {
		return (*elems)[i] < (*elems)[j]
	})
}

type IgnoreSpaces struct {
}

func (s *IgnoreSpaces) Exec(elems *[]string) {
	for i := 0; i < len(*elems); i++ {
		(*elems)[i] = strings.TrimRight((*elems)[i], " ")
	}
}

type IsSorted struct {
}

func (s *IsSorted) Exec(elems *[]string) {
	for i := 1; i < len(*elems); i++ {
		if (*elems)[i-1] > (*elems)[i] {
			log.Println("disorder:", (*elems)[i])
			break
		}
	}
	os.Exit(0)
}
