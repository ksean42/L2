package main

import (
	"os"
	"sort/pkg"
)

func main() {
	if len(os.Args) < 2 {
		pkg.PrintUsage()
		return
	}
	ops, fileName := pkg.ParseFlags()

	strings := pkg.ReadFile(fileName)
	sort := pkg.NewSort(ops, strings)
	sort.ExecOperations()
	pkg.Output(*sort.Elems)
}
