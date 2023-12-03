package day3

import (
	"advent2023/utils"
	"math"
	"regexp"
	"strconv"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 3)
}

func part1(lines []string) any {
	sum := 0

	symbols := []pos{}
	numRanges := []numRange{}

	symbolRegex := regexp.MustCompile("[^0-9.]")
	numberRegex := regexp.MustCompile("[0-9]+")

	for y, line := range lines {
		symbolMatches := symbolRegex.FindAllStringIndex(line, len(line))
		for _, x := range symbolMatches {
			symbols = append(symbols, pos{x: x[0], y: y})
		}
		numberMatches := numberRegex.FindAllStringIndex(line, len(line))
		for _, xRange := range numberMatches {
			num, err := strconv.Atoi(line[xRange[0]:xRange[1]])
			utils.Check(err)
			numRanges = append(numRanges, numRange{
				int:  num,
				y:    y,
				xMin: xRange[0],
				xMax: xRange[1] - 1,
			})
		}
	}

	for _, symbol := range symbols {
		for _, numRange := range numRanges {
			if numRange.isAdjacent(symbol) {
				sum += numRange.int
			}
		}
	}

	return sum
}

func part2(lines []string) any {
	sum := 0

	symbols := []pos{}
	numRanges := []numRange{}

	symbolRegex := regexp.MustCompile("\\*")
	numberRegex := regexp.MustCompile("[0-9]+")

	for y, line := range lines {
		symbolMatches := symbolRegex.FindAllStringIndex(line, len(line))
		for _, x := range symbolMatches {
			symbols = append(symbols, pos{x: x[0], y: y})
		}
		numberMatches := numberRegex.FindAllStringIndex(line, len(line))
		for _, xRange := range numberMatches {
			num, err := strconv.Atoi(line[xRange[0]:xRange[1]])
			utils.Check(err)
			numRanges = append(numRanges, numRange{
				int:  num,
				y:    y,
				xMin: xRange[0],
				xMax: xRange[1] - 1,
			})
		}
	}

	for _, symbol := range symbols {
		adjacentNumbers := []int{}
		for _, numRange := range numRanges {
			if numRange.isAdjacent(symbol) {
				adjacentNumbers = append(adjacentNumbers, numRange.int)
			}
		}
		if len(adjacentNumbers) == 2 {
			sum += adjacentNumbers[0] * adjacentNumbers[1]
		}
	}

	return sum
}

type pos struct {
	x int
	y int
}

type numRange struct {
	int
	y    int
	xMin int
	xMax int
}

func (n numRange) isAdjacent(p pos) bool {
	return math.Abs(float64(n.y-p.y)) <= 1 && n.xMax >= p.x-1 && n.xMin <= p.x+1
}
