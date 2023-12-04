package day4

import (
	"advent2023/utils"
	"math"
	"regexp"
	"slices"
	"strconv"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 4)
}

func part1(lines []string) any {
	sum := 0

	for _, line := range lines {
		card := lineToCard(line)
		sum += card.points()
	}

	return sum
}

func part2(lines []string) any {
	sum := 0
	copies := make([]int, len(lines))

	for i := range lines {
		copies[i] = 1
	}

	for i, line := range lines {
		card := lineToCard(line)
		for x := 1; x <= card.matchingNumbers() && x < len(copies); x++ {
			copies[i+x] += copies[i]
		}
		sum += copies[i]
	}

	return sum
}

func lineToCard(line string) card {
	split := regexp.MustCompile("(:[ ]+)|( \\|[ ]+)").Split(line, len(line))
	winning := regexp.MustCompile("[ ]+").Split(split[1], len(split[1]))
	mine := regexp.MustCompile("[ ]+").Split(split[2], len(split[2]))
	winningNums := make([]int, len(winning))
	myNums := make([]int, len(mine))
	for i, winningNum := range winning {
		intWinningNum, err := strconv.Atoi(winningNum)
		utils.Check(err)
		winningNums[i] = intWinningNum
	}
	for i, myNum := range mine {
		intMyNum, err := strconv.Atoi(myNum)
		utils.Check(err)
		myNums[i] = intMyNum
	}
	return card{
		winningNums,
		myNums,
	}
}

type card struct {
	winningNums []int
	myNums      []int
}

func (c card) matchingNumbers() int {
	matching := 0
	for _, num := range c.myNums {
		if slices.Contains(c.winningNums, num) {
			matching += 1
		}
	}
	return matching
}

func (c card) points() int {
	matching := c.matchingNumbers()
	if matching == 0 {
		return 0
	} else {
		return int(math.Pow(2, float64(matching)-1))
	}
}
