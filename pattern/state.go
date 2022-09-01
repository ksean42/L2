package main

import (
	"fmt"
)

/*
	Состояние — это поведенческий паттерн, позволяющий динамически изменять поведение объекта при смене его состояния.
	Плюсы: код относящийся к одному состоянию сосредоточен в одном месте. Избавляет от лишних условных операторов
	Минусы: Чрезмерно усложняет код, если состояний мало
	В примере используется объект данных , метод поиска по которому меняется в зависимости от его состояния.
*/

type State interface {
	search(needle int)
}

type Sorted struct {
	data myData
}

func (s *Sorted) search(needle int) {
	fmt.Println("Do binary search")
}

type NotSorted struct {
	data myData
}

func (s *NotSorted) search(needle int) {
	fmt.Println("Do naive search")
}

type myData struct {
	data   []int
	sorted State
}

func (m *myData) setState(s State) {
	m.sorted = s
}

func (m *myData) sortIt() {
	fmt.Println("Do sort")
	m.setState(&Sorted{})
}

func (m *myData) add(i int) {
	m.data = append(m.data, 2)

	if isSorted(m.data) {
		m.setState(&Sorted{})
	} else {
		m.setState(&NotSorted{})
	}
}

func (m *myData) doSearch(needle int) {
	m.sorted.search(needle)
}

func isSorted(data []int) bool {
	for i := 1; i < len(data); i++ {
		if data[i-1] > data[i] {
			return false
		}
	}
	return true
}

func newData(data []int) *myData {
	newD := &myData{data: data}
	if isSorted(data) {
		newD.sorted = &Sorted{}
	} else {
		newD.sorted = &NotSorted{}
	}
	return newD
}

func main() {
	data1 := newData([]int{1, 2, 3, 4, 5})
	data2 := newData([]int{3, 4, -1, 2, 1})

	data1.doSearch(1)

	data2.doSearch(1)

	data1.add(2)
	data1.doSearch(1)
}
