package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/moabukar/playground/aoc-2023/utils"
	"github.com/moabukar/playground/aoc-2023/utils/set"
)

//go:embed input.txt
var inputDay string

const (
	EMPTY       = '.'
	UPPER_LEFT  = 'F'
	UPPER_RIGHT = '7'
	LOWER_LEFT  = 'L'
	LOWER_RIGHT = 'J'
	START       = 'S'
	VERTICAL    = '|'
	HORIZONTAL  = '-'
	NORTH       = iota
	SOUTH
	EAST
	WEST
)

func step(grid utils.Grid, pos utils.Pos, from int) (newPos utils.Pos, newFrom int, ok bool) {
	tile, found := grid[pos]

	if !found || tile == EMPTY {
		return pos, from, false
	}
	if tile == START {
		return pos, from, true
	}

	switch tile {
	case VERTICAL:
		if from == NORTH {
			return utils.Pos{X: pos.X, Y: pos.Y + 1}, from, true
		} else if from == SOUTH {
			return utils.Pos{X: pos.X, Y: pos.Y - 1}, from, true
		}
	case HORIZONTAL:
		if from == EAST {
			return utils.Pos{X: pos.X - 1, Y: pos.Y}, from, true
		} else if from == WEST {
			return utils.Pos{X: pos.X + 1, Y: pos.Y}, from, true
		}
	case UPPER_LEFT:
		if from == SOUTH {
			return utils.Pos{X: pos.X + 1, Y: pos.Y}, WEST, true
		} else if from == EAST {
			return utils.Pos{X: pos.X, Y: pos.Y + 1}, NORTH, true
		}
	case UPPER_RIGHT:
		if from == SOUTH {
			return utils.Pos{X: pos.X - 1, Y: pos.Y}, EAST, true
		} else if from == WEST {
			return utils.Pos{X: pos.X, Y: pos.Y + 1}, NORTH, true
		}
	case LOWER_LEFT:
		if from == NORTH {
			return utils.Pos{X: pos.X + 1, Y: pos.Y}, WEST, true
		} else if from == EAST {
			return utils.Pos{X: pos.X, Y: pos.Y - 1}, SOUTH, true
		}
	case LOWER_RIGHT:
		if from == NORTH {
			return utils.Pos{X: pos.X - 1, Y: pos.Y}, EAST, true
		} else if from == WEST {
			return utils.Pos{X: pos.X, Y: pos.Y - 1}, SOUTH, true
		}
	}
	return pos, from, false
}

func findLoop(grid utils.Grid, pos utils.Pos, from int) (path set.Set[utils.Pos], ok bool) {
	path = make(set.Set[utils.Pos])
	var tile, found = grid[pos]
	if !found || tile == EMPTY {
		return path, false
	}

	path.Add(pos)
	for {
		pos, from, found = step(grid, pos, from)
		if !found {
			return path, false
		}
		path.Add(pos)
		tile, _ = grid[pos]
		if tile == START {
			return path, true
		}
	}
}

func findStart(grid utils.Grid) (pos utils.Pos) {
	for pos, tile := range grid {
		if tile == START {
			return pos
		}
	}
	panic("unreachable")
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var grid = utils.BuildGrid(lines)
	var start = findStart(grid)

	var neighbors = []utils.Pos{
		{start.X + 1, start.Y},
		{start.X - 1, start.Y},
		{start.X, start.Y - 1},
		{start.X, start.Y + 1},
	}
	var froms = []int{WEST, EAST, SOUTH, NORTH}

	for i, n := range neighbors {
		loop, found := findLoop(grid, n, froms[i])
		if found {
			var l = len(loop)
			if l%2 == 0 {
				return l / 2
			}
			return l/2 + 1
		}
	}
	panic("no solution found")
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var grid = utils.BuildGrid(lines)
	var start = findStart(grid)

	var neighbors = []utils.Pos{
		{start.X + 1, start.Y},
		{start.X - 1, start.Y},
		{start.X, start.Y - 1},
		{start.X, start.Y + 1},
	}
	var froms = []int{WEST, EAST, SOUTH, NORTH}

	var loopSet set.Set[utils.Pos]
	for i, n := range neighbors {
		var found bool
		loopSet, found = findLoop(grid, n, froms[i])
		if found {
			break
		}
	}

	var res int
	minX, maxX, minY, maxY := utils.GridBounds(grid)
	for y := minY; y <= maxY; y++ {
		var last uint8
		var cpt = 0
		for x := maxX; x >= minX; x-- {
			p := utils.Pos{X: x, Y: y}
			if !loopSet.Contains(p) {
				if cpt%2 == 1 {
					res++
				}
			} else {
				tile, _ := grid[p]
				if last == UPPER_RIGHT && tile == UPPER_LEFT {
					cpt++
				} else if last == UPPER_RIGHT && tile == LOWER_LEFT {
					// do not count
				} else if last == LOWER_RIGHT && tile == LOWER_LEFT {
					cpt++
				} else if last == LOWER_RIGHT && tile == UPPER_LEFT {
					// do not count
				} else if tile != HORIZONTAL {
					last = tile
					cpt++
				}
			}
		}
	}

	return res
}

func main() {
	start := time.Now()
	fmt.Println("Part 1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("Part 2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
