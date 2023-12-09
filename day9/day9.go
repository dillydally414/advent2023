package day9

import (
	"advent2023/utils"
	"regexp"
	"strconv"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 9)
}

func part1(lines []string) any {
	sum := int32(0)

	for _, line := range lines {
		split := regexp.MustCompile("[-0-9]+").FindAllString(line, len(line))
		sequence := make([]int32, len(split))
		for i, s := range split {
			sequence[i] = int32(utils.CheckAndReturn(strconv.Atoi(s)))
		}

		history := generateHistory(sequence)
		extrapolated := extrapolate(history)
		sum += extrapolated
	}

	return sum
}

func part2(lines []string) any {
	sum := int32(0)

	for _, line := range lines {
		split := regexp.MustCompile("[-0-9]+").FindAllString(line, len(line))
		sequence := make([]int32, len(split))
		for i, s := range split {
			sequence[i] = int32(utils.CheckAndReturn(strconv.Atoi(s)))
		}

		history := generateHistory(sequence)
		extrapolated := extrapolateBackwards(history)
		sum += extrapolated
	}

	return sum
}

func generateDiffs(sequence []int32) []int32 {
	diffs := make([]int32, len(sequence)-1)
	for i := range sequence[1:] {
		diffs[i] = sequence[i+1] - sequence[i]
	}
	return diffs
}

func allSame(sequence []int32) bool {
	start := sequence[0]
	for _, num := range sequence {
		if num != start {
			return false
		}
	}
	return true
}

func generateHistory(sequence []int32) [][]int32 {
	history := [][]int32{sequence}
	for !allSame(history[len(history)-1]) {
		history = append(history, generateDiffs(history[len(history)-1]))
	}
	return history
}

func extrapolate(history [][]int32) int32 {
	currDiff := int32(0)
	for i := len(history) - 1; i >= 0; i-- {
		currDiff += history[i][len(history[i])-1]
	}
	return currDiff
}

func extrapolateBackwards(history [][]int32) int32 {
	currDiff := int32(0)
	for i := len(history) - 1; i >= 0; i-- {
		currDiff = history[i][0] - currDiff
	}
	return currDiff
}
