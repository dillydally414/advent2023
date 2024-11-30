package day14

import (
	"advent2023/utils"
	"slices"
	"strings"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 14)
}

func part1(lines []string) any {
	load := 0

	southEdge := len(lines)

	for col := range lines[0] {
		prevRock := -1
		for row := range lines {
			if lines[row][col] == 'O' {
				prevRock++
				load += southEdge - prevRock
			} else if lines[row][col] == '#' {
				prevRock = row
			}
		}
	}

	return load
}

func part2(lines []string) any {
	load := 0

	southEdge := len(lines)

	loadToCycles := make(map[int][]int)

	prevCycle := make([][]rune, len(lines))
	for i, line := range lines {
		prevCycle[i] = []rune(line)
	}

	prevCycles := [][][]rune{prevCycle}
	var loopStart int
	var loopSize int

	for i := 1; i < 1000000000; i++ {
		curr := make([][]rune, len(lines))
		for l, line := range prevCycle {
			curr[l] = slices.Clone(line)
		}
		for direction := 0; direction < 4; direction++ {
			tiltNorth(curr)
			curr = rotate(curr)
		}
		currLoad := 0
		for col := range prevCycle[0] {
			for row := range prevCycle {
				if prevCycle[row][col] == 'O' {
					currLoad += southEdge - row
				}
			}
		}
		loop := false
		if currLoad == 64 {
			println(i)
		}
		for _, idxOfCycle := range loadToCycles[currLoad] {
			if slices.CompareFunc(curr, prevCycles[idxOfCycle], slices.Compare) == 0 {
				loop = true
				loopStart = idxOfCycle
				loopSize = i - idxOfCycle
				break
			}
		}
		if loop {
			break
		}

		loadToCycles[currLoad] = append(loadToCycles[currLoad], i)
		prevCycle = curr
		prevCycles = append(prevCycles, prevCycle)
	}

	cycleToUseIdx := (1000000000-loopStart)%loopSize + loopStart
	cycleToUse := prevCycles[cycleToUseIdx]
	println(loopStart, loopSize, cycleToUseIdx)

	for col := range cycleToUse[0] {
		for row := range cycleToUse {
			if cycleToUse[row][col] == 'O' {
				load += southEdge - row
			}
		}
	}

	return load
}

type coord struct {
	x int
	y int
}

func tiltNorth(platform [][]rune) {
	for col := range platform[0] {
		prevRock := -1
		for row := range platform {
			if platform[row][col] == 'O' {
				prevRock++
				platform[row][col] = '.'
				platform[prevRock][col] = 'O'
			} else if platform[row][col] == '#' {
				prevRock = row
			}
		}
	}
}

func rotate(platform [][]rune) [][]rune {
	newPlatform := make([][]rune, len(platform[0]))
	for i := range platform[0] {
		newPlatform[i] = make([]rune, len(platform))
		for j := range platform {
			newPlatform[i][j] = platform[j][i]
		}
	}
	for _, line := range newPlatform {
		slices.Reverse(line)
	}
	return newPlatform
}

func toString(platform [][]rune) string {
	strs := make([]string, len(platform))
	for i := range platform {
		strs[i] = string(platform[i])
	}
	return strings.Join(strs, "\n")
}
