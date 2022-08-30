package pkg

import (
	"flag"
	"os"
)

// Flag Структура хранения флагов
type Flag struct {
	A int
	B int
	C int
	c bool
	i bool
	v bool
	F bool
	n bool
}

// MyGrep Основной объект программы
type MyGrep struct {
	operation    Operation
	inputStrings *[]string
	pattern      string
	flags        Flag
}

// NewMyGrep Конструктор объекта. Парсит флаги и читает данные из файла в объект
func NewMyGrep() *MyGrep {
	fileName, grep := parseFlags()
	grep.inputStrings = ReadFile(fileName)
	return grep
}

//Do выполняет выбраную стратегию.
func (m *MyGrep) Do() {
	if m.flags.i {
		m.pattern = "(?i)" + m.pattern
	}
	m.operation.Exec(*m.inputStrings, m.pattern)
}

func parseFlags() (string, *MyGrep) {
	flags := Flag{}
	flag.IntVar(&flags.A, "A", 0, "")
	flag.IntVar(&flags.B, "B", 0, "")
	flag.IntVar(&flags.C, "C", 0, "")
	flag.BoolVar(&flags.c, "c", false, "")
	flag.BoolVar(&flags.i, "i", false, "")
	flag.BoolVar(&flags.v, "v", false, "")
	flag.BoolVar(&flags.F, "F", false, "")
	flag.BoolVar(&flags.n, "n", false, "")
	flag.Parse()
	if len(flag.Args()) < 2 {
		os.Exit(1)
	}
	pattern := flag.Args()[0]
	fileName := flag.Args()[1]
	grep := &MyGrep{pattern: pattern, flags: flags}
	if grep.flags.c {
		grep.operation = &Count{}
	} else if grep.flags.v {
		grep.operation = &Invert{}
	} else if grep.flags.C > 0 {
		grep.operation = &Context{After: grep.flags.C, Before: grep.flags.C}
	} else if grep.flags.F {
		grep.operation = &Fixed{}
	} else if grep.flags.n {
		grep.operation = &NumOfLine{}
	} else {
		grep.operation = &Context{After: grep.flags.A, Before: grep.flags.B}
	}
	return fileName, grep
}
