package main

import (
	"flag"
	"fmt"
	"go-text-stats/internal/stats"
	"os"
)

func main() {
	path := flag.String("path", "", "Path to the .txt file")
	flag.Parse()
	if *path == "" {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go -path <file.txt>")
		os.Exit(1)
	}

	file, err := os.Open(*path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Printf("Scanning file: %s\n", *path)

	rows, words, symbols, err := stats.CountTextStats(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error scanning file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Scan finished.\nRows: %d\nWords: %d\nSymbols: %d\n", rows, words, symbols)
}
