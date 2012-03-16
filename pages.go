package npdfpages

import (
	"bufio"
	"io"
	"os"
)

// guaranteed not to overlap or be adjacent
const match = "/Page\x00"

func Pages(reader io.ByteReader) (pages int) {
	i := 0
	for {
		b, err := reader.ReadByte()
		if err != nil {
			return
		}
	check:
		switch match[i] {
		case 0:
			if !(b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z') {
				pages++
			}
			i = 0
			goto check
		case b:
			i++
		default:
			i = 0
		}
	}
	// flow shouldn't reach this point
	return
}

func PagesAtPath(path string) (pages int) {
	if reader, err := os.Open(path); err == nil {
		reader.Chdir()
		pages = Pages(bufio.NewReader(reader))
		reader.Close()
	}
	return
}
