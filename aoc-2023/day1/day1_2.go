package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var numberMap = map[string]int{
	"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4,
	"five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var sum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		first, last := findDigitsWithWords(line)
		if first != -1 && last != -1 {
			value, _ := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
			sum += value
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Sum of calibration values (Part Two):", sum)
}

func findDigitsWithWords(s string) (int, int) {
	var first, last int = -1, -1
	words := regexp.MustCompile(`[a-z]+|\d`).FindAllString(s, -1)
	for _, word := range words {
		if num, exists := numberMap[word]; exists {
			if first == -1 {
				first = num
			}
			last = num
		} else if len(word) == 1 && '0' <= word[0] && word[0] <= '9' {
			num := int(word[0] - '0')
			if first == -1 {
				first = num
			}
			last = num
		}
	}
	return first, last
}
