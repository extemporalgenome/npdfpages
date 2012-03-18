// Copyright 2012 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package npdfpages has no purpose other than to determine the number of pages in
a PDF file. No checking is done to determine that the input is well formatted.

It's possible, but extremely unlikely that PDF streams could contain byte
sequences that result in false positives, though the number returned will be no
less than the number of actual pages in the PDF.
*/
package npdfpages

import (
	"bufio"
	"io"
	"os"
)

const match = "/Page\x00"

// Pages reads the given io.ByteReader until EOF is reached, returning the
// number of pages encountered.
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

// PagesAtPath opens a PDF file at the given file path, returning the number
// of pages found.
func PagesAtPath(path string) (pages int) {
	if reader, err := os.Open(path); err == nil {
		reader.Chdir()
		pages = Pages(bufio.NewReader(reader))
		reader.Close()
	}
	return
}
