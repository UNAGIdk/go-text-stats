package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

// check function shuts down the program if the passed error was not nil
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// CLI-utility that count rows, words and symbols in a .txt file
func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	check(err)
	fmt.Println("Scanning file", filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var rowsAmount, wordsAmount, symbolsAmount int
	for scanner.Scan() {
		rowsAmount++
		wordsAmount += len(strings.Fields(scanner.Text()))
		symbolsAmount += utf8.RuneCountInString(scanner.Text())
	}
	check(scanner.Err())
	fmt.Printf("Scan is finished. Text stats:\nRows amount: %d\nWords amount: %d\nSymbols amount: %d", rowsAmount, wordsAmount, symbolsAmount)
}
