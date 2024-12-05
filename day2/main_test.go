package main

import "testing"

func TestVerifyReportSafety(t *testing.T) {
	tc := []struct {
		line     string
		expected bool
	}{
		{"7 6 4 2 1", true},
		{"1 2 7 8 9", false},
		{"9 7 6 2 1", false},
		{"1 3 2 4 5", false},
		{"8 6 4 4 1", false},
		{"1 3 6 7 9", true},
	}

	for _, tt := range tc {
		tt := tt
		t.Run(tt.line, func(t *testing.T) {
			actual, err := verifyReportSafety(tt.line)
			if actual != tt.expected {
				t.Errorf("Expected %v, but got [%v, %v] for line: %v", tt.expected, actual, err, tt.line)
			}
		})
	}
}

func TestVerifyReportSafetyWithBuffer(t *testing.T) {
	tc := []struct {
		line     string
		expected bool
		err      error
	}{
		{"7 6 4 2 1", true, nil},
		{"1 2 7 8 9", false, nil},
		{"9 7 6 2 1", false, nil},
		{"1 3 2 4 5", true, nil},
		{"8 6 4 4 1", true, nil},
		{"1 3 6 7 9", true, nil},
		{"24 25 28 31 28", true, nil},
		{"41 44 45 48 49 50 50", true, nil},
		{"5 8 10 13 15 16 17 21", true, nil},
		{"11 13 16 17 19 26", true, nil},
		{"79 81 78 79 82 84", false, nil},
		{"16 19 20 18 20 22 25 22", false, nil},
		{"84 87 90 92 94 97 96 96", false, nil},
		{"86 87 88 91 88 91 95", false, nil},
		{"8 10 10 11 13", true, nil},

		{"84 87 90 92 94 96 96", true, nil},
		{"84 87 90 92 94 96 101", true, nil},

		{"80 87 90 92 94 96", true, nil},
		{"80 87 90 92 94 96 101", false, nil},

		{"40 43 41 44 49", false, nil},
		{"40 44 41 40 37", true, nil},
	}

	for _, tt := range tc {
		tt := tt
		t.Run(tt.line, func(t *testing.T) {
			actual, err := verifyReportSafetyWithBuffer(tt.line)
			if actual != tt.expected {
				t.Errorf("Expected %v, but got [%v, [%v,%v]] for line: %v", tt.expected, actual, err[0], err[1], tt.line)
			}
		})
	}
}
