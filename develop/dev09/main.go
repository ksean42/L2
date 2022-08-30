package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please enter valid URL: scheme://www.example.com")
	}

	URL, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Get(URL.String())
	if err != nil {
		log.Fatal(err)
	}

	filename := strings.Split(URL.Host, ".")
	file, err := os.Create(filename[1] + ".html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range body {
		fmt.Fprint(file, string(v))
	}

}
