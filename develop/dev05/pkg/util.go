package pkg

import (
	"bufio"
	"log"
	"os"
)

// ReadFile Функция чтения данных из файла
func ReadFile(fileName string) *[]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var strings []string
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		strings = append(strings, sc.Text())
	}
	return &strings
}
