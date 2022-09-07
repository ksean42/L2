package server

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseConfig() string {
	file, err := os.Open("config")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	read := bufio.NewScanner(file)
	for read.Scan() {
		str := strings.Split(read.Text(), "=")
		if str[0] == "port" {
			_, err := strconv.Atoi(str[1])
			if err != nil {
				log.Fatal("Incorrect config")
			}
			return ":" + str[1]
		}
	}
	return ""
}
