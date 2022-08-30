package main

import (
	"anagram/pkg"
)

func main() {
	dict := pkg.ReadDict()
	pkg.PrintResult(pkg.SearchAnagrams(dict))
}
