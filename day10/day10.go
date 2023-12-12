package day10

import (
	"advent2023/utils"
	"fmt"
	"slices"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 10)
}

func part1(lines []string) any {
	pipes := make(map[coord][]coord)
	var startLoc coord

	for y, line := range lines {
		for x, char := range line {
			currCord := coord{x, y}
			switch char {
			case '|':
				pipes[currCord] = []coord{{x: x, y: y - 1}, {x: x, y: y + 1}}
				break
			case '-':
				pipes[currCord] = []coord{{x: x - 1, y: y}, {x: x + 1, y: y}}
				break
			case 'L':
				pipes[currCord] = []coord{{x: x, y: y - 1}, {x: x + 1, y: y}}
				break
			case 'J':
				pipes[currCord] = []coord{{x: x - 1, y: y}, {x: x, y: y - 1}}
				break
			case '7':
				pipes[currCord] = []coord{{x: x - 1, y: y}, {x: x, y: y + 1}}
				break
			case 'F':
				pipes[currCord] = []coord{{x: x, y: y + 1}, {x: x + 1, y: y}}
				break
			case '.':
				pipes[currCord] = []coord{}
				break
			case 'S':
				pipes[currCord] = []coord{{x: x, y: y + 1}, {x: x + 1, y: y}, {x: x, y: y - 1}, {x: x - 1, y: y}}
				startLoc = currCord
				break
			default:
				panic(fmt.Sprintln("Spot not valid: %r", char))
			}
		}
	}

	var currLoc coord
	prevLoc := startLoc
	neighbors := pipes[startLoc]

	for _, neighbor := range neighbors {
		if slices.Contains(pipes[neighbor], startLoc) {
			currLoc = neighbor
			break
		}
	}

	length := 1

	for currLoc != startLoc {
		length++
		nextChoices := pipes[currLoc]
		if prevLoc == nextChoices[0] {
			prevLoc = currLoc
			currLoc = nextChoices[1]
		} else {
			prevLoc = currLoc
			currLoc = nextChoices[0]
		}
	}

	return length / 2
}

func part2(lines []string) any {
	pipes := make(map[coord][]coord)
	var startLoc coord

	for y, line := range lines {
		for x, char := range line {
			currCord := coord{x, y}
			switch char {
			case '|':
				pipes[currCord] = []coord{{x: x, y: y - 1}, {x: x, y: y + 1}}
				break
			case '-':
				pipes[currCord] = []coord{{x: x - 1, y: y}, {x: x + 1, y: y}}
				break
			case 'L':
				pipes[currCord] = []coord{{x: x, y: y - 1}, {x: x + 1, y: y}}
				break
			case 'J':
				pipes[currCord] = []coord{{x: x - 1, y: y}, {x: x, y: y - 1}}
				break
			case '7':
				pipes[currCord] = []coord{{x: x - 1, y: y}, {x: x, y: y + 1}}
				break
			case 'F':
				pipes[currCord] = []coord{{x: x, y: y + 1}, {x: x + 1, y: y}}
				break
			case '.':
				pipes[currCord] = []coord{}
				break
			case 'S':
				pipes[currCord] = []coord{{x: x, y: y + 1}, {x: x + 1, y: y}, {x: x, y: y - 1}, {x: x - 1, y: y}}
				startLoc = currCord
				break
			default:
				panic(fmt.Sprintln("Spot not valid: %r", char))
			}
		}
	}

	var currLoc coord
	prevLoc := startLoc
	neighbors := pipes[startLoc]

	for _, neighbor := range neighbors {
		if slices.Contains(pipes[neighbor], startLoc) {
			currLoc = neighbor
			break
		}
	}

	pipesInLoop := make(map[coord]bool)
	pipesInLoop[startLoc] = true

	for currLoc != startLoc {
		pipesInLoop[currLoc] = true
		nextChoices := pipes[currLoc]
		if prevLoc == nextChoices[0] {
			prevLoc = currLoc
			currLoc = nextChoices[1]
		} else {
			prevLoc = currLoc
			currLoc = nextChoices[0]
		}
	}

	insideCount := 0

	for y, line := range lines {
		inside := false
		var prevPipe rune
		for x := range line {
			if pipesInLoop[coord{x, y}] {
				if line[x] == '|' || (line[x] == 'J' && prevPipe == 'F') || (line[x] == '7' && prevPipe == 'L') {
					inside = !inside
				}
				if line[x] != '-' {
					prevPipe = rune(line[x])
				}
			} else if inside {
				insideCount++
			}
		}
	}

	return insideCount
}

type coord struct {
	x int
	y int
}
