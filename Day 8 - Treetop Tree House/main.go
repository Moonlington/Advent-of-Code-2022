package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AOCForest struct {
	trees [][]int
}

func newForest() *AOCForest {
	return &AOCForest{
		trees: [][]int{},
	}
}

func (f *AOCForest) addTreesFromLine(line string) {
	treeStrings := strings.Split(line, "")
	var addToForest []int
	for _, tree := range treeStrings {
		treeInt, err := strconv.Atoi(tree)
		if err != nil {
			panic(err)
		}
		addToForest = append(addToForest, treeInt)
	}
	f.trees = append(f.trees, addToForest)
}

func (f *AOCForest) printForest() {
	for _, y := range f.trees {
		for _, x := range y {
			fmt.Print(x)
		}
		fmt.Println()
	}
}

func (f *AOCForest) getTreeHeight(x, y int) int {
	if x < 0 || y < 0 {
		return -1
	}
	if x >= len(f.trees[0]) || y >= len(f.trees) {
		return -1
	}
	return f.trees[y][x]
}

func (f *AOCForest) isTreeVisible(x, y int) bool {
	height := f.getTreeHeight(x, y)
	blockedByRight := false
	blockedByLeft := false
	blockedByUp := false
	blockedByDown := false
	for dist := 1; ; dist++ {
		if f.getTreeHeight(x+dist, y) >= height {
			blockedByRight = true
		}
		if f.getTreeHeight(x-dist, y) >= height {
			blockedByLeft = true
		}
		if f.getTreeHeight(x, y+dist) >= height {
			blockedByDown = true
		}
		if f.getTreeHeight(x, y-dist) >= height {
			blockedByUp = true
		}
		if f.getTreeHeight(x+dist, y) == -1 && f.getTreeHeight(x-dist, y) == -1 && f.getTreeHeight(x, y+dist) == -1 && f.getTreeHeight(x, y-dist) == -1 {
			break
		}
	}
	return !(blockedByDown && blockedByUp && blockedByLeft && blockedByRight)
}

func (f *AOCForest) getAmountVisibleTrees() int {
	sum := 0
	for y := range f.trees {
		for x := range f.trees[y] {
			if f.isTreeVisible(x, y) {
				sum++
			}
		}
	}
	return sum
}

func (f *AOCForest) getTreeScenicScore(x, y int) int {
	height := f.getTreeHeight(x, y)
	score := 1
	dist := 1
	for {
		if f.getTreeHeight(x+dist, y) >= height {
			score *= dist
			break
		}
		if f.getTreeHeight(x+dist, y) == -1 {
			score *= dist - 1
			break
		}
		dist++
	}
	dist = 1
	for {
		if f.getTreeHeight(x-dist, y) >= height {
			score *= dist
			break
		}
		if f.getTreeHeight(x-dist, y) == -1 {
			score *= dist - 1
			break
		}
		dist++
	}
	dist = 1
	for {
		if f.getTreeHeight(x, y+dist) >= height {
			score *= dist
			break
		}
		if f.getTreeHeight(x, y+dist) == -1 {
			score *= dist - 1
			break
		}
		dist++
	}
	dist = 1
	for {
		if f.getTreeHeight(x, y-dist) >= height {
			score *= dist
			break
		}
		if f.getTreeHeight(x, y-dist) == -1 {
			score *= dist - 1
			break
		}
		dist++
	}
	return score
}

func (f *AOCForest) getMaxScenicScore() int {
	max := 0
	for y := range f.trees {
		for x := range f.trees[y] {
			score := f.getTreeScenicScore(x, y)
			if score > max {
				max = score
			}
		}
	}
	return max
}

var Forest *AOCForest = newForest()

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		Forest.addTreesFromLine(scanner.Text())
	}
	fmt.Println(Forest.getAmountVisibleTrees()) // Part 1 Solution
	fmt.Println(Forest.getMaxScenicScore())     // Part 2 Solution
}
