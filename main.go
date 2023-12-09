package main

import (
	"advent2023/day1"
	"advent2023/day2"
	"advent2023/day3"
	"advent2023/day4"
	"advent2023/day5"
	"advent2023/day6"
	"advent2023/day7"
	"advent2023/day8"
	"advent2023/utils"
	"flag"
	"fmt"
)

var days = []utils.Day{
	day1.GetDay(),
	day2.GetDay(),
	day3.GetDay(),
	day4.GetDay(),
	day5.GetDay(),
	day6.GetDay(),
	day7.GetDay(),
	day8.GetDay(),
}

func getDay(day int) utils.Day {
	if day < 1 || day > len(days) {
		panic(fmt.Sprintf("No implementation found for day %d", day))
	}

	return days[day-1]
}

func getRunner() utils.DayRunner {
	dayPtr := flag.Int("day", -1, "day to run solution on")
	actualPtr := flag.Bool("actual", false, "run on actual input")
	part2Ptr := flag.Bool("part2", false, "run part 2")
	flag.Parse()

	day := getDay(*dayPtr)
	return utils.DayRunner{
		Day:       day,
		RunActual: *actualPtr,
		RunPart2:  *part2Ptr,
	}
}

func main() {
	runner := getRunner()
	runner.Run()
}
