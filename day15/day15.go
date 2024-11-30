package day15

import (
	"advent2023/utils"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 15)
}

func part1(lines []string) any {
	sum := 0

	for _, step := range strings.Split(lines[0], ",") {
		sum += hash(step)
	}

	return sum
}

func part2(lines []string) any {
	boxes := make([]box, 256)

	for i := range boxes {
		boxes[i] = box{
			lensLength: make(map[string]int),
			lensOrder:  []string{},
		}
	}

	for _, step := range strings.Split(lines[0], ",") {
		label := regexp.MustCompile("[a-z]+").FindString(step)
		boxNum := hash(label)
		if regexp.MustCompile("-").MatchString(step) {
			boxes[boxNum].removeLens(label)
		} else {
			focalLength := utils.CheckAndReturn(strconv.Atoi(regexp.MustCompile("[1-9]").FindString(step)))
			boxes[boxNum].addLens(label, focalLength)
		}
	}

	power := 0

	for boxNum, box := range boxes {
		power += box.power(boxNum)
	}

	return power
}

func hash(s string) int {
	value := 0
	for _, r := range s {
		value += int(r)
		value *= 17
		value = value % 256
	}
	return value
}

type box struct {
	lensOrder  []string
	lensLength map[string]int
}

func (b *box) addLens(label string, focalLength int) {
	b.lensLength[label] = focalLength
	idxOf := slices.Index(b.lensOrder, label)
	if idxOf == -1 {
		b.lensOrder = append(b.lensOrder, label)
	}
}

func (b *box) removeLens(label string) {
	delete(b.lensLength, label)
	idxOf := slices.Index(b.lensOrder, label)
	if idxOf != -1 {
		b.lensOrder = slices.Delete(b.lensOrder, idxOf, idxOf+1)
	}
}

func (b box) power(boxNum int) int {
	power := 0
	for i, lensLabel := range b.lensOrder {
		power += (boxNum + 1) * (i + 1) * b.lensLength[lensLabel]
	}
	return power
}
