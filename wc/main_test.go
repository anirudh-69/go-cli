package main

import (
	"bytes"
	"io"
	"testing"
)

// TestCountLines tests the count function set to count lines
func TestCountLines(t *testing.T) {
	tests := []struct {
		name     string
		input    io.Reader
		expected int
	}{
		{"Multiple lines", bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1"), 3},
		{"Single line", bytes.NewBufferString("word1 word2 word3"), 1},
		{"Empty input", bytes.NewBufferString(""), 0},
		{"Nil Input", nil, 0},
		{"Trailing newline", bytes.NewBufferString("word1 word2 word3\n"), 1},
		{"Multiple newlines", bytes.NewBufferString("word1 word2 word3\n\n\n"), 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.input
			result := count(b, true)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d instead.\n", tt.expected, result)
			}
		})
	}
}

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    io.Reader
		expected int
	}{
		{"Multiple words", bytes.NewBufferString("word1 word2 word3 word4\n"), 4},
		{"Single word", bytes.NewBufferString("word1"), 1},
		{"Empty input", bytes.NewBufferString(""), 0},
		{"Multiple spaces", bytes.NewBufferString("word1   word2  word3"), 3},
		{"Mixed spaces and newlines", bytes.NewBufferString("word1 word2\nword3 word4"), 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.input
			result := count(b, false)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d instead.\n", tt.expected, result)
			}
		})
	}
}
