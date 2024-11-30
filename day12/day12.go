package day12

import (
	"advent2023/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 12)
}

func part1(lines []string) any {
	arrangements := 0

	for _, line := range lines {
		split := strings.Split(line, " ")
		damagedCounts := regexp.MustCompile("[0-9]+").FindAllString(split[1], len(split[1]))
		damagedCountsNum := make([]int, len(damagedCounts))
		brokenSum := 0
		for i, num := range damagedCounts {
			damagedCountsNum[i] = utils.CheckAndReturn(strconv.Atoi(num))
			brokenSum += damagedCountsNum[i]
		}
		matchingRegex := createSpringRegex(damagedCountsNum)
		knownBroken := len(regexp.MustCompile("[#]").FindAllString(split[0], len(split[0])))
		missingBroken := brokenSum - knownBroken
		unknownLocs := regexp.MustCompile("[?]").FindAllStringIndex(split[0], len(split[0]))
		possibleIndices := combin.Combinations(len(unknownLocs), missingBroken)
		for _, p := range possibleIndices {
			newStrRunes := []rune(strings.Clone(split[0]))
			for _, replacement := range p {
				newStrRunes[unknownLocs[replacement][0]] = '#'
			}
			newStr := strings.ReplaceAll(string(newStrRunes), "?", ".")
			if matchingRegex.MatchString(newStr) {
				arrangements++
			}
		}
	}

	return arrangements
}

func part2(lines []string) any {
	arrangements := 0

	for _, line := range lines {
		split := strings.Split(line, " ")
		newSpringsSlice := make([]string, 5)
		for i := range newSpringsSlice {
			newSpringsSlice[i] = split[0]
		}
		newSpringsStr := strings.Join(newSpringsSlice, "?")
		damagedCounts := regexp.MustCompile("[0-9]+").FindAllString(split[1], len(split[1]))
		damagedCountsNum := make([]int, len(damagedCounts)*len(newSpringsSlice))

		brokenSum := 0
		for i, num := range damagedCounts {
			numVal := utils.CheckAndReturn(strconv.Atoi(num))
			for mult := range newSpringsSlice {
				damagedCountsNum[i+len(damagedCounts)*mult] = numVal
			}
			brokenSum += damagedCountsNum[i] * len(newSpringsSlice)
		}

		matchingRegex := createSpringRegex(damagedCountsNum)
		knownBroken := len(regexp.MustCompile("[#]").FindAllString(newSpringsStr, len(newSpringsStr)))
		missingBroken := brokenSum - knownBroken
		unknownLocs := regexp.MustCompile("[?]").FindAllStringIndex(newSpringsStr, len(newSpringsStr))
		possibleIndices := combin.Combinations(len(unknownLocs), missingBroken)
		fmt.Println(line, missingBroken, len(possibleIndices))
		currArrangements := 0
		for _, p := range possibleIndices {
			newStrRunes := []rune(strings.Clone(newSpringsStr))
			for _, replacement := range p {
				newStrRunes[unknownLocs[replacement][0]] = '#'
			}
			newStr := strings.ReplaceAll(string(newStrRunes), "?", ".")
			if matchingRegex.MatchString(newStr) {
				currArrangements++
			}
		}
		arrangements += currArrangements
	}

	return arrangements
}

type coord struct {
	x int
	y int
}

func createSpringRegex(groups []int) regexp.Regexp {
	groupStrings := make([]string, len(groups))
	for i, g := range groups {
		groupStrings[i] = fmt.Sprintf("#{%d}", g)
	}
	return *regexp.MustCompile("\\.*" + strings.Join(groupStrings, "\\.+") + "\\.*")
}
