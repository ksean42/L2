package main

import (
	"mygrep/pkg"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	grep := pkg.NewMyGrep()
	grep.Do()
}
