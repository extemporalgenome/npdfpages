package npdfpages

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

var exp, _ = regexp.Compile(`/Page\b`)

func PagesInRuneReader(reader io.RuneReader) (pages int) {
	for exp.FindReaderIndex(reader) != nil {
		pages++
	}
	return
}

func PagesAtPath(path string) (pages int) {
	if reader, err := os.Open(path); err == nil {
		pages = PagesInRuneReader(bufio.NewReader(reader))
		reader.Close()
	}
	return
}
