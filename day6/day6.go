package day6

import (
	"advent2023/utils"
	"regexp"
	"strconv"
	"strings"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 6)
}

func part1(lines []string) any {
	times := regexp.MustCompile("[0-9]+").FindAllString(lines[0], len(lines[0]))
	distances := regexp.MustCompile("[0-9]+").FindAllString(lines[1], len(lines[1]))
	waysToWin := 1

	for i := range times {
		ms := utils.CheckAndReturn(strconv.Atoi(times[i]))
		mm := utils.CheckAndReturn(strconv.Atoi(distances[i]))
		currMsRange := []int{0, ms / 2}
		for currMsRange[0] <= currMsRange[1] {
			targetMs := (currMsRange[0] + currMsRange[1]) / 2
			distanceTraveled := (targetMs) * (ms - targetMs)
			distanceTraveledMinus1 := (targetMs - 1) * (ms - (targetMs - 1))
			if distanceTraveled <= mm {
				currMsRange[0] = targetMs + 1
			} else if distanceTraveledMinus1 <= mm {
				waysToWin *= (ms - 2*targetMs) + 1
				break
			} else {
				currMsRange[1] = targetMs - 1
			}
		}
	}

	return waysToWin
}

func part2(lines []string) any {
	times := regexp.MustCompile("[0-9]+").FindAllString(lines[0], len(lines[0]))
	distances := regexp.MustCompile("[0-9]+").FindAllString(lines[1], len(lines[1]))

	ms := utils.CheckAndReturn(strconv.Atoi(strings.Join(times, "")))
	mm := utils.CheckAndReturn(strconv.Atoi(strings.Join(distances, "")))

	currMsRange := []int{0, ms / 2}
	for currMsRange[0] <= currMsRange[1] {
		targetMs := (currMsRange[0] + currMsRange[1]) / 2
		distanceTraveled := (targetMs) * (ms - targetMs)
		distanceTraveledMinus1 := (targetMs - 1) * (ms - (targetMs - 1))
		if distanceTraveled <= mm {
			currMsRange[0] = targetMs + 1
		} else if distanceTraveledMinus1 <= mm {
			return (ms - 2*targetMs) + 1
		} else {
			currMsRange[1] = targetMs - 1
		}
	}

	return nil
}
