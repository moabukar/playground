package main

import (
	_ "embed"
	"fmt"

	// "github.com/moabukar/playground/aoc-2023/utils"
	"strings"
)

//go:embed input.txt
var inputDay string

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}
func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

func (s Set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Intersect(other Set[T]) Set[T] {
	res := NewSet[T]()
	if s.Len() < other.Len() {
		for elem := range s {
			if other.Contains(elem) {
				res.Add(elem)
			}
		}
	} else {
		for elem := range other {
			if s.Contains(elem) {
				res.Add(elem)
			}
		}
	}
	return res
}

////////////////////////////////////////

func winning(line string) int {
	split := func(c rune) bool { return c == ':' || c == '|' }
	fields := strings.FieldsFunc(line, split)

	var winningNumbers = NewSet[string]()
	for _, n := range strings.Fields(fields[1]) {
		winningNumbers.Add(n)
	}
	var numbers = NewSet[string]()
	for _, n := range strings.Fields(fields[2]) {
		numbers.Add(n)
	}
	return winningNumbers.Intersect(numbers).Len()
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	var res int
	for _, line := range lines {
		v := winning(line)
		if v > 0 {
			res += 1 << (v - 1)
		}
	}
	return res
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	var cards []int
	var matchs []int
	for _, line := range lines {
		v := winning(line)
		cards = append(cards, 1)
		matchs = append(matchs, v)
	}
	for i := 0; i < len(cards); i++ {
		if matchs[i] > 0 {
			for j := i + 1; j < i+1+matchs[i]; j++ {
				cards[j] += cards[i]
			}
		}
	}
	var res int
	for _, c := range cards {
		res += c
	}
	return res
}

func main() {
	fmt.Println("Part 1: ", Part1(inputDay))

	fmt.Println("Part 2: ", Part2(inputDay))
}
