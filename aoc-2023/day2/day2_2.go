package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	games, err := readGamesFromFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return
	}

	totalPower := 0
	for _, game := range games {
		_, minRed, minGreen, minBlue := findMinimumCubes(game)
		power := minRed * minGreen * minBlue
		totalPower += power
	}

	fmt.Printf("Total power of the minimum sets: %d\n", totalPower)
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

func findMinimumCubes(game string) (int, int, int, int) {
	id, rounds := parseGame(game)
	minRed, minGreen, minBlue := 0, 0, 0

	for _, round := range rounds {
		minRed = max(minRed, round["red"])
		minGreen = max(minGreen, round["green"])
		minBlue = max(minBlue, round["blue"])
	}

	return id, minRed, minGreen, minBlue
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
