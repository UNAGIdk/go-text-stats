package stats

import (
	"bufio"
	"io"
	"strings"
	"unicode/utf8"
)

// CountTextStats reads from the provided reader and returns
// amount of rows, words and Unicode symbols it contains
func CountTextStats(r io.Reader) (rows, words, symbols int, err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		rows++
		words += len(strings.Fields(line))
		symbols += utf8.RuneCountInString(line)
	}

	err = scanner.Err()
	if err != nil {
		return rows, words, symbols, err
	}

	return rows, words, symbols, nil
}
