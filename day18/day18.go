package day18

import (
	"advent2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 18)
}

func part1(lines []string) any {
	trench := make(map[coord]bool)
	curr := coord{
		x: 0,
		y: 0,
	}
	minX, maxX, minY, maxY := 0, 0, 0, 0

	for _, command := range lines {
		split := strings.Split(command, " ")
		direction, length := rune(split[0][0]), utils.CheckAndReturn(strconv.Atoi(split[1]))
		delta := coord{x: 0, y: 0}
		switch direction {
		case 'U':
			delta.y = -1
			break
		case 'D':
			delta.y = 1
			break
		case 'R':
			delta.x = 1
			break
		case 'L':
			delta.x = -1
			break
		default:
			panic(fmt.Sprintln("unrecognized symbol", direction))
		}
		for i := 1; i <= length; i++ {
			c := coord{x: curr.x + delta.x*i, y: curr.y + delta.y*i}
			trench[c] = true
		}
		curr.x += delta.x * length
		curr.y += delta.y * length
		minX, maxX, minY, maxY = min(minX, curr.x), max(maxX, curr.x), min(minY, curr.y), max(maxY, curr.y)
	}

	insideCount := 0

	for y := minY; y <= maxY; y++ {
		inside := false
		prevTrench := -1
		for x := minX; x <= maxX; x++ {
			if trench[coord{x, y}] {
				if !trench[coord{x: x - 1, y: y}] || (trench[coord{x: prevTrench, y: y + 1}] && trench[coord{x: x, y: y + 1}]) || (trench[coord{x: prevTrench, y: y - 1}] && trench[coord{x: x, y: y - 1}]) {
					inside = !inside
				}
				if !trench[coord{x: x - 1, y: y}] {
					prevTrench = x
				}
			} else if inside {
				insideCount++
			}
		}
	}

	return insideCount + len(trench)
}

func part2(lines []string) any {
	trench := make(map[coord]bool)
	curr := coord{
		x: 0,
		y: 0,
	}
	minX, maxX, minY, maxY := 0, 0, 0, 0

	for _, command := range lines {
		split := strings.Split(command, " ")
		length := int(utils.CheckAndReturn(strconv.ParseInt(string(split[2][2:7]), 16, 0)))
		direction := utils.CheckAndReturn(strconv.Atoi(string(split[2][7])))
		delta := coord{x: 0, y: 0}
		switch direction {
		case 3:
			delta.y = -1
			break
		case 1:
			delta.y = 1
			break
		case 0:
			delta.x = 1
			break
		case 2:
			delta.x = -1
			break
		default:
			panic(fmt.Sprintln("unrecognized symbol", direction))
		}
		for i := 1; i <= length; i++ {
			c := coord{x: curr.x + delta.x*i, y: curr.y + delta.y*i}
			trench[c] = true
		}
		curr.x += delta.x * length
		curr.y += delta.y * length
		minX, maxX, minY, maxY = min(minX, curr.x), max(maxX, curr.x), min(minY, curr.y), max(maxY, curr.y)
	}

	insideCount := 0

	for y := minY; y <= maxY; y++ {
		inside := false
		prevTrench := -1
		for x := minX; x <= maxX; x++ {
			if trench[coord{x, y}] {
				if !trench[coord{x: x - 1, y: y}] || (trench[coord{x: prevTrench, y: y + 1}] && trench[coord{x: x, y: y + 1}]) || (trench[coord{x: prevTrench, y: y - 1}] && trench[coord{x: x, y: y - 1}]) {
					inside = !inside
				}
				if !trench[coord{x: x - 1, y: y}] {
					prevTrench = x
				}
			} else if inside {
				insideCount++
			}
		}
	}

	return insideCount + len(trench)
}

type coord struct {
	x int
	y int
}
