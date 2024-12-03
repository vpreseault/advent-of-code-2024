package main

import (
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
		if verifySafety(line) {
			safeCounter++
		}

	}

	fmt.Println("Safe levels: ", safeCounter)
}

func verifySafety(line string) bool {
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

		if one == two {
			return false
		}

		diff := one - two
		if i == 0 {
			if diff < 0 {
				increasing = true
			}
		} else {
			if (diff < 0 && !increasing) || (diff > 0 && increasing) {
				return false
			}
		}

		if abs(diff) > 3 {
			return false
		}
	}
	return true
}

func part2(data []byte) {

}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
