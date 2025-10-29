package stats

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
		fmt.Fprintln(os.Stderr)
		os.Exit(1)
	}

	return rows, words, symbols, nil
}
