package day11

import (
	"advent2023/utils"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 11)
}

func part1(lines []string) any {
	galaxies := []coord{}

	doubledRows := make([]bool, len(lines))
	doubledColumns := make([]bool, len(lines[0]))

	for x := range lines[0] {
		doubledColumns[x] = true
	}

	for y, line := range lines {
		doubledRows[y] = true
		for x, char := range line {
			if char == '#' {
				galaxies = append(galaxies, coord{x, y})
				doubledRows[y] = false
				doubledColumns[x] = false
			}
		}
	}

	sum := 0

	for i, start := range galaxies {
		for _, end := range galaxies[i+1:] {
			xLow, xHigh := min(start.x, end.x), max(start.x, end.x)
			yLow, yHigh := min(start.y, end.y), max(start.y, end.y)
			for x := xLow; x < xHigh; x++ {
				sum++
				if doubledColumns[x] {
					sum++
				}
			}
			for y := yLow; y < yHigh; y++ {
				sum++
				if doubledRows[y] {
					sum++
				}
			}
		}
	}

	return sum
}

func part2(lines []string) any {
	galaxies := []coord{}

	doubledRows := make([]bool, len(lines))
	doubledColumns := make([]bool, len(lines[0]))

	for x := range lines[0] {
		doubledColumns[x] = true
	}

	for y, line := range lines {
		doubledRows[y] = true
		for x, char := range line {
			if char == '#' {
				galaxies = append(galaxies, coord{x, y})
				doubledRows[y] = false
				doubledColumns[x] = false
			}
		}
	}

	sum := 0

	for i, start := range galaxies {
		for _, end := range galaxies[i+1:] {
			xLow, xHigh := min(start.x, end.x), max(start.x, end.x)
			yLow, yHigh := min(start.y, end.y), max(start.y, end.y)
			for x := xLow; x < xHigh; x++ {
				sum++
				if doubledColumns[x] {
					sum += 999999
				}
			}
			for y := yLow; y < yHigh; y++ {
				sum++
				if doubledRows[y] {
					sum += 999999
				}
			}
		}
	}

	return sum
}

type coord struct {
	x int
	y int
}
