package day8

import (
	"advent2023/utils"
	"regexp"
	"slices"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 8)
}

func part1(lines []string) any {
	lr := lines[0]
	directions := make(map[string]instruction)

	for _, line := range lines[2:] {
		split := regexp.MustCompile("[A-Z]{3}").FindAllString(line, 3)
		directions[split[0]] = instruction{left: split[1], right: split[2]}
	}

	currLoc := "AAA"
	currIdx := 0

	for currLoc != "ZZZ" {
		if string(lr[currIdx%len(lr)]) == "L" {
			currLoc = directions[currLoc].left
		} else {
			currLoc = directions[currLoc].right
		}
		currIdx++
	}

	return currIdx
}

func part2(lines []string) any {
	lr := lines[0]
	directions := make(map[string]instruction)

	startingLocs := []string{}
	allLocs := make([]string, len(lines)-2)

	for i, line := range lines[2:] {
		split := regexp.MustCompile("[0-9|A-Z]{3}").FindAllString(line, len(line))
		directions[split[0]] = instruction{left: split[1], right: split[2]}
		allLocs[i] = split[0]
		if string(split[0][2]) == "A" {
			startingLocs = append(startingLocs, split[0])
		}
	}

	oneCycle := make(map[string][]string)
	for _, loc := range allLocs {
		oneCycle[loc] = make([]string, len(lr))
	}

	currLocs := slices.Clone(allLocs)

	for step := 0; step < len(lr); step++ {
		if string(lr[step]) == "L" {
			for i, loc := range currLocs {
				nextStep := directions[loc].left
				oneCycle[allLocs[i]][step] = nextStep
				currLocs[i] = nextStep
			}
		} else {
			for i, loc := range currLocs {
				nextStep := directions[loc].right
				oneCycle[allLocs[i]][step] = nextStep
				currLocs[i] = nextStep
			}
		}
	}

	cycleDurations := make(map[string]cycle)
	currLocs = startingLocs[:]
	cyclesSoFar := make(map[string][]string)

	for i, loc := range startingLocs {
		cyclesSoFar[startingLocs[i]] = []string{}
		currLoc := loc
		for cycleCt := 0; cycleCt >= 0; cycleCt++ {
			currLoc = oneCycle[currLoc][len(lr)-1]
			prevIdx := slices.Index(cyclesSoFar[startingLocs[i]], currLoc)
			if prevIdx != -1 {
				cycleDurations[loc] = cycle{
					start:  prevIdx + 1,
					length: cycleCt - prevIdx,
				}
				break
			}
			cyclesSoFar[startingLocs[i]] = append(cyclesSoFar[startingLocs[i]], currLoc)
		}
	}

	completeLoop := len(lr)
	for _, loc := range startingLocs {
		completeLoop *= cycleDurations[loc].length
	}

	return completeLoop
}

type instruction struct {
	left  string
	right string
}

func stepsEndingWithZ(cycles [][]string) []bool {
	result := make([]bool, len(cycles[0]))
	for i := range result {
		result[i] = true
		for _, cycle := range cycles {
			if string(cycle[i][2]) != "Z" {
				result[i] = false
				break
			}
		}
	}

	return result
}

type cycle struct {
	start  int
	length int
}
