package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Action int

const (
	Rock Action = iota + 1
	Paper
	Scissors
)

type WinCondition int

const (
	Lose WinCondition = iota * 3
	Draw
	Win
)

func parseAction(input string) Action {
	if input == "A" || input == "X" {
		return Rock
	} else if input == "B" || input == "Y" {
		return Paper
	} else if input == "C" || input == "Z" {
		return Scissors
	} else {
		panic(input)
	}
}

func parseWinCondition(input string) WinCondition {
	switch input {
	case "X":
		return Lose
	case "Y":
		return Draw
	case "Z":
		return Win
	}
	panic(input)
}

func calculateRound(opponent Action, yours Action) int {
	if opponent == yours {
		return 3
	}
	switch opponent {
	case Rock:
		if yours == Paper {
			return 6
		} else {
			return 0
		}
	case Paper:
		if yours == Scissors {
			return 6
		} else {
			return 0
		}
	case Scissors:
		if yours == Rock {
			return 6
		} else {
			return 0
		}
	}
	panic(opponent)
}

func getActionToGetWinCondition(opponent Action, winCondition WinCondition) Action {
	if winCondition == Draw {
		return opponent
	}
	switch opponent {
	case Rock:
		if winCondition == Win {
			return Paper
		} else {
			return Scissors
		}
	case Paper:
		if winCondition == Win {
			return Scissors
		} else {
			return Rock
		}
	case Scissors:
		if winCondition == Win {
			return Rock
		} else {
			return Paper
		}
	}
	panic(opponent)
}

func calculateScore(input string) int {
	score := 0
	left, right, _ := strings.Cut(input, " ")
	opponentAction := parseAction(left)
	yourAction := parseAction(right)
	score += int(yourAction)
	score += calculateRound(opponentAction, yourAction)
	return score
}

func calculateScore2(input string) int {
	score := 0
	left, right, _ := strings.Cut(input, " ")
	opponentAction := parseAction(left)
	WinCondition := parseWinCondition(right)
	score += int(WinCondition)
	score += int(getActionToGetWinCondition(opponentAction, WinCondition))
	return score
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	score := 0
	score2 := 0
	for scanner.Scan() {
		score += calculateScore(scanner.Text())
		score2 += calculateScore2(scanner.Text())
	}
	fmt.Println(score)  // Part 1 Answer
	fmt.Println(score2) // Part 2 Answer
}
