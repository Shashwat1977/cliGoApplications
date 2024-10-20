package main

import (
	"strings"
	"testing"
)

func Test_countWords(t *testing.T) {
	reader := strings.NewReader("Test string to check count of words")
	res := count(reader, false)
	expected := 7

	if res != expected {
		t.Errorf("Expected %v, Got: %v", expected, res)
	}
}

func Test_wc(t *testing.T) {
	tableTest := []struct {
		name     string
		inp      string
		expected int
		flag     bool
	}{
		{"count words", "Test string to check count of words", 7, false},
		{"count lines", "Test string to check count of lines", 1, true},
		{"count multiple lines", "First line.\nSecond line.", 2, true},
	}

	for _, tt := range tableTest {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.inp)
			res := count(reader, tt.flag)
			if res != tt.expected {
				t.Errorf("Expected %v, Got: %v", tt.expected, res)
			}
		})
	}
}
