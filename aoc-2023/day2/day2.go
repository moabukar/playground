package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	totalRed := 12
	totalGreen := 13
	totalBlue := 14

	games, err := readGamesFromFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return
	}

	sumOfPossibleGameIDs := solve(games, totalRed, totalGreen, totalBlue)
	fmt.Printf("Sum of possible game IDs: %d\n", sumOfPossibleGameIDs)
}

func readGamesFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var games []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		games = append(games, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return games, nil
}

func solve(games []string, totalRed, totalGreen, totalBlue int) int {
	sum := 0

	for _, game := range games {
		id, rounds := parseGame(game)
		if isGamePossible(rounds, totalRed, totalGreen, totalBlue) {
			sum += id
		}
	}

	return sum
}

func parseGame(game string) (int, []map[string]int) {
	var id int
	fmt.Sscanf(game, "Game %d:", &id)

	roundsData := strings.Split(game[strings.Index(game, ":")+1:], ";")
	rounds := make([]map[string]int, len(roundsData))

	for i, round := range roundsData {
		rounds[i] = make(map[string]int)
		colors := strings.Split(strings.TrimSpace(round), ", ")

		for _, color := range colors {
			var count int
			var colorName string
			fmt.Sscanf(color, "%d %s", &count, &colorName)
			rounds[i][colorName] = count
		}
	}

	return id, rounds
}

func isGamePossible(rounds []map[string]int, totalRed, totalGreen, totalBlue int) bool {
	maxRed, maxGreen, maxBlue := 0, 0, 0

	for _, round := range rounds {
		maxRed = max(maxRed, round["red"])
		maxGreen = max(maxGreen, round["green"])
		maxBlue = max(maxBlue, round["blue"])
	}

	return maxRed <= totalRed && maxGreen <= totalGreen && maxBlue <= totalBlue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
