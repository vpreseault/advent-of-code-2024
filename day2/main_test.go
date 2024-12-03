package main

import "testing"

func TestVerifySafety(t *testing.T) {
	tc := []struct {
		line     string
		expected bool
	}{
		{
			line:     "24 25 28 31 28",
			expected: false,
		},
		{
			line:     "41 44 45 48 49 50 50",
			expected: false,
		},
		{
			line:     "5 8 10 13 15 16 17 21",
			expected: false,
		},
		{
			line:     "11 13 16 17 19 26",
			expected: false,
		},
		{
			line:     "79 81 78 79 82 84",
			expected: false,
		},
		{
			line:     "16 19 20 18 20 22 25 22",
			expected: false,
		},
		{
			line:     "84 87 90 92 94 97 96 96",
			expected: false,
		},
		{
			line:     "86 87 88 91 88 91 95",
			expected: false,
		},
		{
			line:     "40 43 41 44 49",
			expected: false,
		},
		{
			line:     "8 10 10 11 13",
			expected: false,
		},
		{
			line:     "91 94 95 95 92",
			expected: false,
		},
		{
			line:     "18 19 20 21 23 25",
			expected: true,
		},
		{
			line:     "15 17 19 20 23",
			expected: true,
		},
		{
			line:     "35 37 37 39 40 43 50",
			expected: false,
		},
		{
			line:     "2 5 6 10 12",
			expected: false,
		},
	}

	for _, tt := range tc {
		tt := tt
		t.Run(tt.line, func(t *testing.T) {
			actual := verifySafety(tt.line)
			if actual != tt.expected {
				t.Errorf("Expected %v, but got %v for line: %v", tt.expected, actual, tt.line)
			}
		})
	}
}
