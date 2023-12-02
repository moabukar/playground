package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	spelledNumbers := map[string]string{
		"one": "1", "two": "2", "three": "3", "four": "4",
		"five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
	}

	scanner := bufio.NewScanner(os.Stdin)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.FieldsFunc(line, func(r rune) bool {
			return r < 'a' || r > 'z'
		})

		first, last := "", ""
		for _, word := range words {
			if val, exists := spelledNumbers[word]; exists {
				if first == "" {
					first = val
				}
				last = val
			} else if _, err := strconv.Atoi(word); err == nil && len(word) == 1 {
				if first == "" {
					first = word
				}
				last = word
			}
		}

		if first != "" && last != "" {
			number, _ := strconv.Atoi(first + last)
			sum += number
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println(sum)
}
