package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1(data)
	part2(data)
}

func part1(data []byte) {
	lines := strings.Split(string(data), "\n")
	leftList, rightList := []int{}, []int{}
	for _, line := range lines {
		distances := strings.Split(line, "   ")
		if leftDistance, err := strconv.Atoi(distances[0]); err == nil {
			leftList = append(leftList, leftDistance)
		}
		if rightDistance, err := strconv.Atoi(distances[1]); err == nil {
			rightList = append(rightList, rightDistance)
		}
	}

	leftList = quicksort(leftList)
	rightList = quicksort(rightList)

	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		distance := leftList[i] - rightList[i]
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance
	}

	fmt.Printf("Total distance: %v\n", totalDistance)
}

func part2(data []byte) {
	lines := strings.Split(string(data), "\n")
	leftList, rightList := []int{}, make(map[int]int)
	for _, line := range lines {
		distances := strings.Split(line, "   ")
		if leftDistance, err := strconv.Atoi(distances[0]); err == nil {
			leftList = append(leftList, leftDistance)
		}

		if rightDistance, err := strconv.Atoi(distances[1]); err == nil {
			if val, ok := rightList[rightDistance]; ok {
				rightList[rightDistance] = val + 1
			} else {
				rightList[rightDistance] = 1
			}
		}
	}

	similarity := 0
	for i := 0; i < len(leftList); i++ {
		if val, ok := rightList[leftList[i]]; ok {
			similarity += leftList[i] * val
		}
	}

	fmt.Printf("Similarity: %v\n", similarity)
}

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	pivot := len(arr) / 2

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quicksort(arr[:left])
	quicksort(arr[left+1:])

	return arr
}
