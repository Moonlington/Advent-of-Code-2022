package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	sumContains := 0
	sumOverlap := 0
	for scanner.Scan() {
		first, second, _ := strings.Cut(scanner.Text(), ",")
		if AssignmentContains(first, second) || AssignmentContains(second, first) {
			sumContains++
		}
		if AssignmentOverlap(first, second) || AssignmentOverlap(second, first) {
			sumOverlap++
		}
	}
	fmt.Println(sumContains) // Part 1 Solution
	fmt.Println(sumOverlap)  // Part 2 Solution
}

func AssignmentContains(container, substring string) bool {
	leftBound, rightBound, _ := strings.Cut(container, "-")
	leftBoundInt, _ := strconv.Atoi(leftBound)
	rightBoundInt, _ := strconv.Atoi(rightBound)
	leftSubstring, rightSubstring, _ := strings.Cut(substring, "-")
	leftSubstringInt, _ := strconv.Atoi(leftSubstring)
	rightSubstringInt, _ := strconv.Atoi(rightSubstring)

	return leftSubstringInt >= leftBoundInt && rightSubstringInt <= rightBoundInt
}

func AssignmentOverlap(container, substring string) bool {
	leftBound, rightBound, _ := strings.Cut(container, "-")
	leftBoundInt, _ := strconv.Atoi(leftBound)
	rightBoundInt, _ := strconv.Atoi(rightBound)
	leftSubstring, rightSubstring, _ := strings.Cut(substring, "-")
	leftSubstringInt, _ := strconv.Atoi(leftSubstring)
	rightSubstringInt, _ := strconv.Atoi(rightSubstring)

	return leftSubstringInt <= rightBoundInt && rightSubstringInt >= leftBoundInt
}
