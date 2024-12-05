package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1(data)
	part2(data)
}

func part1(data []byte) {
	safeCounter := 0
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if val, err := verifyReportSafety(line); err == nil && val {
			safeCounter++
		}

	}

	fmt.Println("Safe levels: ", safeCounter)
}

func verifyReportSafety(line string) (bool, error) {
	increasing := false
	entries := strings.Split(line, " ")
	for i := 0; i < len(entries)-1; i++ {
		one, err := strconv.Atoi(entries[i])
		if err != nil {
			log.Fatal(err)
		}

		two, err := strconv.Atoi(entries[i+1])
		if err != nil {
			log.Fatal(err)
		}

		if val, err := verifyLevelSafety(i, &increasing, one, two); !val {
			return false, err
		}
	}
	return true, nil
}

func part2(data []byte) {
	safeCounter := 0
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if val, err := verifyReportSafetyWithBuffer(line); err == nil && val {
			safeCounter++
		}

	}

	fmt.Println("Safe levels with buffer: ", safeCounter)
}

func verifyReportSafetyWithBuffer(line string) (bool, []error) {
	increasing := false
	bufferUsed := false
	var errs []error
	entries := strings.Split(line, " ")
	for i := 0; i < len(entries)-1; i++ {
		one, err := strconv.Atoi(entries[i])
		if err != nil {
			log.Fatal(err)
		}

		two, err := strconv.Atoi(entries[i+1])
		if err != nil {
			log.Fatal(err)
		}

		if val, err := verifyLevelSafety(i, &increasing, one, two); !val {
			errs = append(errs, err)
			if i+2 < len(entries) && !bufferUsed {
				three, err := strconv.Atoi(entries[i+2])
				if err != nil {
					log.Fatal(err)
				}

				bufferUsed = true
				if val, err := verifyLevelSafety(i, &increasing, one, three); !val {
					errs = append(errs, err)
					return false, errs
				}
				i++ // skip two
			}
		}

	}
	return true, errs
}

func verifyLevelSafety(index int, increasing *bool, n1, n2 int) (bool, error) {
	if n1 == n2 {
		return false, errors.New("duplicate values")
	}

	diff := n1 - n2
	if index == 0 {
		if diff < 0 {
			*increasing = true
		}
	} else {
		if (diff < 0 && !*increasing) || (diff > 0 && *increasing) {
			if *increasing {
				return false, fmt.Errorf("wanted to increase, but decreased")
			} else {
				return false, fmt.Errorf("wanted to decrease, but increased")
			}

		}
	}

	if abs(diff) > 3 {
		return false, fmt.Errorf("difference is greater than 3")
	}
	return true, nil
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
