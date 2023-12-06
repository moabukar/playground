package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputDay string

func buildArray(line string) []int {
	var res []int
	for _, field := range strings.Fields(line) {
		if v, err := strconv.Atoi(strings.TrimSpace(field)); err == nil {
			res = append(res, v)
		}
	}
	return res
}

func concatBuild(line string) int {
	s := ""
	for _, field := range strings.Fields(line) {
		v := strings.TrimSpace(field)
		if _, err := strconv.Atoi(v); err == nil {
			s = s + v
		}
	}
	res, _ := strconv.Atoi(s)
	return res
}

func calculateDistance(hold, maxTime int) int {
	return (maxTime - hold) * hold
}

func solve(time, dist int) int {
	res := 0
	for i := 0; i < time; i++ {
		d := calculateDistance(i, time)
		if d > dist {
			res++
		}
	}
	return res
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var times = buildArray(lines[0])
	var dist = buildArray(lines[1])

	res := 1
	for i, t := range times {
		res *= solve(t, dist[i])
	}
	return res
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	return solve(concatBuild(lines[0]), concatBuild(lines[1]))
}

func main() {
	fmt.Println("Part 1: ", Part1(inputDay))

	fmt.Println("Part 2: ", Part2(inputDay))
}
