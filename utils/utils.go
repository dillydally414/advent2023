package utils

import (
	"bufio"
	"fmt"
	"os"
)

// helper function to panic on error
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

type Part func(lines []string) (output any)

type Day struct {
	part1        Part
	part2        Part
	actualInput  string
	sampleInput  string
	sampleInput2 string
}

func NewDay(part1 Part, part2 Part, dayCount int) Day {
	dayPrefix := fmt.Sprintf("day%d/", dayCount)
	actualInput := dayPrefix + "actual_input.txt"
	sampleInput := dayPrefix + "sample_input.txt"
	sampleInput2 := dayPrefix + "sample_input_2.txt"

	day := Day{
		part1,
		part2,
		actualInput,
		sampleInput,
		sampleInput2,
	}

	if _, err := os.Stat(sampleInput2); os.IsNotExist(err) {
		day.sampleInput2 = sampleInput
	}

	return day
}

func (day Day) determineFileName(actual bool, part2 bool) string {
	if actual {
		return day.actualInput
	} else {
		if part2 {
			return day.sampleInput2
		} else {
			return day.sampleInput
		}
	}
}

type DayRunner struct {
	Day
	RunActual bool
	RunPart2  bool
}

// abstract runner that passes line-split input into part1 and part2 handlers, with inputs
// determined by command-line flags
func (runner DayRunner) Run() {
	f, err := os.Open(runner.determineFileName(runner.RunActual, runner.RunPart2))
	Check(err)

	scanner := bufio.NewScanner(f)

	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	var output any

	if runner.RunPart2 {
		output = runner.part2(lines)
	} else {
		output = runner.part1(lines)
	}

	fmt.Println(output)
}
