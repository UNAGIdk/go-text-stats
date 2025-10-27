package stats

import (
	"strings"
	"testing"
)

// TestCountTextStats verifies that CountTextStats correctly counts
// lines, words, and Unicode characters in a mixed language input.
func TestCountTextStats(t *testing.T) {
	input := "Hello world\nПривет Go"
	r := strings.NewReader(input)

	rows, words, symbols, err := CountTextStats(r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rows != 2 {
		t.Errorf("expected 2 rows, got %d", rows)
	}

	if words != 4 {
		t.Errorf("expected 4 words, got %d", words)
	}

	if symbols != 20 {
		t.Errorf("expected 20 symbols, got %d", symbols)
	}
}

// TestEmptyInput ensures that CountTextStats returns zeros
// when given an empty input.
func TestEmptyInput(t *testing.T) {
	r := strings.NewReader("")

	rows, words, symbols, err := CountTextStats(r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rows != 0 || words != 0 || symbols != 0 {
		t.Errorf("expected all zeros, got rows=%d, words=%d, symbols=%d", rows, words, symbols)
	}
}

// TestSingleLine checks that CountTextStats handles
// a single-line input correctly.
func TestSingleLine(t *testing.T) {
	r := strings.NewReader("One line text")

	rows, words, symbols, err := CountTextStats(r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rows != 1 {
		t.Errorf("expected 1 row, got %d", rows)
	}
	if words != 3 {
		t.Errorf("expected 3 words, got %d", words)
	}
	if symbols != 13 {
		t.Errorf("expected 13 symbols, got %d", symbols)
	}
}
