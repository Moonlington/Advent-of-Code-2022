package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(inputFile)
	sumPriority := 0
	for scanner.Scan() {
		sacks := scanner.Text()
		sumPriority += itemToPriority(getCommonItem(sacks[:len(sacks)/2], sacks[len(sacks)/2:]))
	}
	fmt.Println(sumPriority) // Part 1 Solution
	inputFile.Close()

	inputFile2, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile2.Close()
	scanner2 := bufio.NewScanner(inputFile2)
	sumPriority2 := 0
	for scanner2.Scan() {
		sack1 := scanner2.Text()
		scanner2.Scan()
		sack2 := scanner2.Text()
		scanner2.Scan()
		sack3 := scanner2.Text()
		sumPriority2 += itemToPriority(getCommonItem2(sack1, sack2, sack3))
	}
	fmt.Println(sumPriority2) // Part 2 Solution
}

func itemToPriority(item string) int {
	return strings.Index(alphabet, item) + 1
}

func getCommonItem(first string, second string) string {
	for _, item := range first {
		for _, item2 := range second {
			if item == item2 {
				return string(item)
			}
		}
	}
	panic("Nothing Common!")
}

func getCommonItem2(first string, second string, third string) string {
	for _, item := range first {
		for _, item2 := range second {
			for _, item3 := range third {
				if item == item2 && item2 == item3 {
					return string(item)
				}
			}
		}
	}
	panic("Nothing Common!")
}
