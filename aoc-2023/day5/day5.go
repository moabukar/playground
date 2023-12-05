package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type SeedMap struct {
	off       string
	to        string
	mapValues []mapValue
}

type mapValue struct {
	source    int
	maxSource int
	target    int
}

func NewSeedMap(off string, to string) *SeedMap {
	return &SeedMap{
		off: off,
		to:  to,
	}
}

func (sm *SeedMap) Add(target int, source int, offset int) {
	sm.mapValues = append(sm.mapValues, mapValue{
		source:    source,
		maxSource: source + offset,
		target:    target,
	})
}

func (sm *SeedMap) GetDestination(given int) int {
	for _, mv := range sm.mapValues {
		if mv.source <= given && given < mv.maxSource {
			return mv.target + (given - mv.source)
		}
	}
	return given
}

func main() {
	startTime := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	seedStrs := strings.Split(scanner.Text(), ": ")[1]
	seeds := convertToIntSlice(strings.Split(seedStrs, " "))

	var seedMaps []*SeedMap
	var currentMap *SeedMap

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.HasSuffix(line, "map:") {
			parts := strings.Split(line, " ")[0]
			offTo := strings.Split(parts, "-to-")
			currentMap = NewSeedMap(offTo[0], offTo[1])
			seedMaps = append(seedMaps, currentMap)
		} else {
			nums := convertToIntSlice(strings.Split(line, " "))
			currentMap.Add(nums[0], nums[1], nums[2])
		}
	}

	minLocation := int(^uint(0) >> 1) // Max int
	for _, seed := range seeds {
		for _, sm := range seedMaps {
			seed = sm.GetDestination(seed)
		}
		if seed < minLocation {
			minLocation = seed
		}
	}

	fmt.Println(minLocation)
	fmt.Printf("Time: %.6f seconds\n", time.Since(startTime).Seconds())
}

func convertToIntSlice(strs []string) []int {
	var ints []int
	for _, s := range strs {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}
