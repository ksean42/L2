package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type cut struct {
	input     string
	field     int
	delimiter string
	separated bool
}

func (c *cut) out() {
	input := strings.Split(c.input, c.delimiter)
	if len(input) == 1 {
		if c.separated {
			fmt.Println()
			return
		}
		fmt.Println(input[0])
	}
	if c.field-1 < len(input) {
		fmt.Println(input[c.field-1])
	} else {
		fmt.Println()
	}

}

func parseFlags(cut *cut) {
	flag.IntVar(&cut.field, "f", 0, "-f choose fields")
	flag.StringVar(&cut.delimiter, "d", "\t", "-d choose delimiter")
	flag.BoolVar(&cut.separated, "s", false, "-s print strings only with delimiter")
	flag.Parse()
	if cut.field == 0 || cut.delimiter == "" {
		log.Fatal("Invalid parameters")
	}
}

func main() {
	cut := &cut{}
	parseFlags(cut)
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		cut.input = sc.Text()
		cut.out()
	}
}
