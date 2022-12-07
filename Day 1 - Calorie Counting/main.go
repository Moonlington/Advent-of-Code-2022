package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	caloriesPerElf := []int{}
	i := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			i++;
			continue
		}
		calorie, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		caloriesPerElf[i] += 
	}
	for _, e := range strings.Split(inputString, "\r\n\r\n") {
		calories := 0
		for _, c := range strings.Split(e, "\r\n") {
			if c == "" {
				continue
			}
			calorie, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			calories += calorie
		}
		caloriesPerElf = append(caloriesPerElf, calories)
	}
	sort.Slice(caloriesPerElf, func(i, j int) bool {
		return caloriesPerElf[i] > caloriesPerElf[j]
	})
	fmt.Println(caloriesPerElf[0])                                         // Part 1 answer
	fmt.Println(caloriesPerElf[0] + caloriesPerElf[1] + caloriesPerElf[2]) // Part 2 answer
}
