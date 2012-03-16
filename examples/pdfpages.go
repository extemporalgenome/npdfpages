package main

import (
	"fmt"
	"github.com/extemporalgenome/npdfpages"
	"os"
)

const PAD_DIGITS = 4

func main() {
	exitcode := 0
	format := fmt.Sprint("%", PAD_DIGITS, "d %s\n")
	for _, path := range os.Args[1:] {
		pages := npdfpages.PagesAtPath(path)
		if pages == 0 {
			exitcode = 1
		}
		fmt.Printf(format, pages, path)
	}
	os.Exit(exitcode)
}
