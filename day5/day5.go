package day5

import (
	"advent2023/utils"
	"cmp"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 5)
}

func part1(lines []string) any {
	seeds := regexp.MustCompile("[0-9]+").FindAllString(lines[0], len(lines[0]))
	seedNums := make([]int, len(seeds))
	for i, str := range seeds {
		num, err := strconv.Atoi(str)
		utils.Check(err)
		seedNums[i] = num
	}

	currMapLines := []string{}
	maps := []almanacMap{}

	for _, line := range lines[2:] {
		if len(line) == 0 {
			maps = append(maps, linesToMap(currMapLines))
			currMapLines = []string{}
		} else {
			currMapLines = append(currMapLines, line)
		}
	}
	maps = append(maps, linesToMap(currMapLines))

	currValues := seedNums

	for _, m := range maps {
		for i, value := range currValues {
			currValues[i] = m.convert(value)
		}
	}

	return slices.Min(currValues)
}

func part2(lines []string) any {
	seeds := regexp.MustCompile("[0-9]+ [0-9]+").FindAllString(lines[0], len(lines[0]))
	seedRanges := make([]numRange, len(seeds))
	for i, str := range seeds {
		seedRangeSplit := strings.Split(str, " ")
		seedRanges[i] = numRange{
			start:  utils.CheckAndReturn(strconv.Atoi(seedRangeSplit[0])),
			length: utils.CheckAndReturn(strconv.Atoi(seedRangeSplit[1])),
		}
	}

	currMapLines := []string{}
	maps := []almanacMap{}

	for _, line := range lines[2:] {
		if len(line) == 0 {
			maps = append(maps, linesToMap(currMapLines))
			currMapLines = []string{}
		} else {
			currMapLines = append(currMapLines, line)
		}
	}
	maps = append(maps, linesToMap(currMapLines))

	currRanges := seedRanges

	for _, m := range maps {
		newCurrRanges := []numRange{}
		for _, sourceRange := range currRanges {
			for _, destRange := range m.convertRange(sourceRange) {
				newCurrRanges = append(newCurrRanges, destRange)
			}
		}
		currRanges = newCurrRanges
	}

	return slices.MinFunc(currRanges, func(range1 numRange, range2 numRange) int { return cmp.Compare(range1.start, range2.start) }).start
}

type numRange struct {
	start  int
	length int
}

func (n numRange) overlap(n2 numRange) (included []numRange, excluded []numRange) {
	if n.start+n.length < n2.start || n2.start+n2.length < n.start {
		// no overlap
		return []numRange{}, []numRange{n}
	} else {
		included := numRange{start: max(n.start, n2.start), length: min(n.start+n.length, n2.start+n2.length) - max(n.start, n2.start)}
		// some or all of range is included
		excludedStart := numRange{start: n.start, length: n2.start - n.start}
		excludedEnd := numRange{start: n2.start + n2.length, length: n.start + n.length - (n2.start + n2.length)}
		excluded := []numRange{}
		if excludedStart.length > 0 {
			excluded = append(excluded, excludedStart)
		}
		if excludedEnd.length > 0 {
			excluded = append(excluded, excludedEnd)
		}
		return []numRange{included}, excluded
	}
}

type almanacMap struct {
	sourceType string
	destType   string
	mappings   []almanacMapEntry
}

type almanacMapEntry struct {
	destStart   int
	sourceStart int
	rangeLength int
}

func linesToMap(lines []string) almanacMap {
	firstLine := regexp.MustCompile("-| ").Split(lines[0], len(lines[0]))
	sourceType, destType := firstLine[0], firstLine[2]
	mappings := make([]almanacMapEntry, len(lines)-1)

	for i, line := range lines[1:] {
		splitLine := strings.Split(line, " ")
		mappings[i] = almanacMapEntry{
			destStart:   utils.CheckAndReturn(strconv.Atoi(splitLine[0])),
			sourceStart: utils.CheckAndReturn(strconv.Atoi(splitLine[1])),
			rangeLength: utils.CheckAndReturn(strconv.Atoi(splitLine[2])),
		}
	}

	return almanacMap{
		sourceType,
		destType,
		mappings,
	}
}

func (a almanacMap) convert(source int) int {
	for _, mapping := range a.mappings {
		if source >= mapping.sourceStart && source < mapping.sourceStart+mapping.rangeLength {
			return mapping.destStart + source - mapping.sourceStart
		}
	}
	return source
}

func (a almanacMap) convertRange(sourceRange numRange) []numRange {
	destRanges := []numRange{}
	remainingRanges := []numRange{sourceRange}
	for _, mapping := range a.mappings {
		newRemainingRanges := []numRange{}
		for _, remainingRange := range remainingRanges {
			included, excluded := remainingRange.overlap(numRange{start: mapping.sourceStart, length: mapping.rangeLength})
			newRemainingRanges = append(newRemainingRanges, excluded...)
			destIncluded := make([]numRange, len(included))
			for i, includedRange := range included {
				destIncluded[i] = numRange{start: includedRange.start + mapping.destStart - mapping.sourceStart, length: includedRange.length}
			}
			destRanges = append(destRanges, destIncluded...)
		}
		remainingRanges = newRemainingRanges
	}
	return append(destRanges, remainingRanges...)
}
