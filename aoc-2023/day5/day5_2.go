package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type SeedRange struct {
	start int
	end   int
}

type SeedRangeList []SeedRange

func NewSeedMap(off string, to string) *SeedMap {
	return &SeedMap{
		off: off,
		to:  to,
	}
}

func (sm *SeedMap) Add(target, source, offset int) {
	sm.mapValues = append(sm.mapValues, mapValue{source, source + offset, target})
}

func (sm *SeedMap) Sort() {
	sort.Slice(sm.mapValues, func(i, j int) bool {
		return sm.mapValues[i].source < sm.mapValues[j].source
	})
}

func (sm *SeedMap) GetDestination(ranges SeedRangeList) SeedRangeList {
	var newRanges SeedRangeList
	for _, r := range ranges {
		start, end := r.start, r.end
		foundMapping := false
		for _, mv := range sm.mapValues {
			if start < mv.source {
				if end < mv.source {
					newRanges = append(newRanges, SeedRange{start, end})
					foundMapping = true
					break
				}
				if end <= mv.maxSource {
					newRanges = append(newRanges, SeedRange{start, mv.source})
					newRanges = append(newRanges, SeedRange{mv.target, mv.target + (end - mv.source)})
					foundMapping = true
					break
				}
				newRanges = append(newRanges, SeedRange{start, mv.source})
				newRanges = append(newRanges, SeedRange{mv.target, mv.target + (mv.maxSource - mv.source)})
				start = mv.maxSource
			} else if start < mv.maxSource {
				if end <= mv.maxSource {
					newRanges = append(newRanges, SeedRange{mv.target + (start - mv.source), mv.target + (end - mv.source)})
					foundMapping = true
					break
				}
				newRanges = append(newRanges, SeedRange{mv.target + (start - mv.source), mv.target + (mv.maxSource - mv.source)})
				start = mv.maxSource
			}
		}
		if !foundMapping {
			newRanges = append(newRanges, SeedRange{start, end})
		}
	}
	return newRanges
}

func (srl *SeedRangeList) Add(start, end int) {
	for i, r := range *srl {
		if start < r.start {
			if end < r.start {
				*srl = append((*srl)[:i], append(SeedRangeList{{start, end}}, (*srl)[i:]...)...)
				return
			}
			if end <= r.end {
				*srl = append((*srl)[:i], append(SeedRangeList{{start, r.start}}, (*srl)[i:]...)...)
				return
			}
			*srl = append((*srl)[:i], append(SeedRangeList{{start, r.start}}, (*srl)[i:]...)...)
			start = r.end
		} else if start < r.end {
			if end <= r.end {
				return
			}
			start = r.end
		}
	}
	*srl = append(*srl, SeedRange{start, end})
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

	seedRangeStrs := strings.Split(scanner.Text(), ": ")[1]
	seedStrs := strings.Split(seedRangeStrs, " ")
	var seeds SeedRangeList
	for i := 0; i < len(seedStrs); i += 2 {
		start, _ := strconv.Atoi(seedStrs[i])
		offset, _ := strconv.Atoi(seedStrs[i+1])
		seeds.Add(start, start+offset)
	}

	var seedMaps []*SeedMap
	var currentMap *SeedMap

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.HasSuffix(line, "map:") {
			offTo := strings.Split(strings.Split(line, " ")[0], "-to-")
			currentMap = NewSeedMap(offTo[0], offTo[1])
			seedMaps = append(seedMaps, currentMap)
		} else {
			values := strings.Split(line, " ")
			target, _ := strconv.Atoi(values[0])
			source, _ := strconv.Atoi(values[1])
			offset, _ := strconv.Atoi(values[2])
			currentMap.Add(target, source, offset)
		}
	}

	for _, sm := range seedMaps {
		sm.Sort()
		seeds = sm.GetDestination(seeds)
	}

	minLocation := int(^uint(0) >> 1) // Max int
	for _, r := range seeds {
		if r.start < minLocation {
			minLocation = r.start
		}
	}

	fmt.Println(minLocation)
	fmt.Printf("Time: %.6f seconds\n", time.Since(startTime).Seconds())
}
