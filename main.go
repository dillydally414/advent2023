package main

import (
	"advent2023/day1"
	"advent2023/day10"
	"advent2023/day11"
	"advent2023/day12"
	"advent2023/day13"
	"advent2023/day14"
	"advent2023/day15"
	"advent2023/day16"
	"advent2023/day17"
	"advent2023/day18"
	"advent2023/day19"
	"advent2023/day2"
	"advent2023/day20"
	"advent2023/day21"
	"advent2023/day22"
	"advent2023/day23"
	"advent2023/day24"
	"advent2023/day25"
	"advent2023/day3"
	"advent2023/day4"
	"advent2023/day5"
	"advent2023/day6"
	"advent2023/day7"
	"advent2023/day8"
	"advent2023/day9"
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
	day9.GetDay(),
	day10.GetDay(),
	day11.GetDay(),
	day12.GetDay(),
	day13.GetDay(),
	day14.GetDay(),
	day15.GetDay(),
	day16.GetDay(),
	day17.GetDay(),
	day18.GetDay(),
	day19.GetDay(),
	day20.GetDay(),
	day21.GetDay(),
	day22.GetDay(),
	day23.GetDay(),
	day24.GetDay(),
	day25.GetDay(),
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
