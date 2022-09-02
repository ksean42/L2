package pkg

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

// PrintUsage выводит usage
func PrintUsage() {
	fmt.Printf("Usage: [-k *num*] [-n] [-r] [-u] [-b] [-c] [filename]\n" +
		"-k *num* is for sort by column number (delimeter is space by default)\n" +
		"-n is for sort by numeric value\n" +
		"-r is for sort in reverse order\n" +
		"-u to delete non unique strings\n" +
		"-b to ignore tail spaces\n" +
		"-c to check if data is sorted\n")
}

// ReadFile читает файл в массив строк
func ReadFile(fileName string) *[]string {
	var strings []string
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		strings = append(strings, sc.Text())
	}
	return &strings
}

// Output вывод в файл
func Output(strings []string) {
	file, err := os.Create("sorted.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, value := range strings {
		fmt.Fprintln(file, value)
	}
}

// ParseFlags парсит флаги
func ParseFlags() ([]Operation, string) {
	var (
		k        int
		n        bool
		r        bool
		u        bool
		b        bool
		c        bool
		fileName string
	)

	flag.IntVar(&k, "k", 0, "-k *num* is for sort by column number (delimeter is space by default)")
	flag.BoolVar(&n, "n", false, "-n is for sort by numeric value")
	flag.BoolVar(&r, "r", false, "-r is for sort in reverse order")
	flag.BoolVar(&u, "u", false, "-u to delete duplicate strings")
	flag.BoolVar(&b, "b", false, "-b to ignore tail spaces")
	flag.BoolVar(&c, "c", false, "-c to check if data is sorted")
	flag.Parse()
	if len(flag.Args()) == 0 {
		PrintUsage()
		os.Exit(1)
	}
	fileName = flag.Args()[0]
	ops := make([]Operation, 0, 6)
	if c { // если есть флаг с то остальные флаги игнорируются
		ops = append(ops, &IsSorted{})
		return ops, fileName
	}
	ops = append(ops, &DefaultSort{})
	if b {
		ops = append(ops, &IgnoreSpaces{})
	}
	if k > 0 {
		ops = append(ops, &SortByColumn{Column: k - 1})
	}
	if n {
		ops = append(ops, &SortByNum{})
	}
	if r {
		ops = append(ops, &ReverseSort{})
	}
	if u {
		ops = append(ops, &Unique{})
	}
	return ops, fileName
}
