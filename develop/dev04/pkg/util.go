package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadDict() *[]string {
	var dict []string
	file, err := os.Open("dictionary.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		dict = append(dict, sc.Text())
	}
	return &dict
}
func PrintResult(res map[string][]string) {
	for k, v := range res {
		fmt.Print(k, ": ")
		fmt.Println(v)
	}
}
