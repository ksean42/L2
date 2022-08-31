package main

import "fmt"

/*
	Команда - поведенческий паттерн, позволяющий превращать запросы в объекты и передавать их как аргументы,
	ставить запросы в очередь, логировать их, а также поддерживать отмену операций.

*/

type Command interface {
	exec(data []int)
}

type Sorter struct {
	data    []int
	command Command
}

func (s Sorter) do() {
	s.command.exec(s.data)
}

type quickSort struct {
}

func (q quickSort) exec(data []int) {
	fmt.Println("data is sorted by quick sort!")
}

type bubbleSort struct {
}

func (b bubbleSort) exec(data []int) {
	fmt.Println("data is sorted by bubble sort!")
}

func main() {
	sorter := &Sorter{data: make([]int, 0, 10)}
	sorter.command = quickSort{}
	sorter.do()

	sorter.command = bubbleSort{}
	sorter.do()
}
