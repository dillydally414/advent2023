package day2

import (
	"advent2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 2)
}

func part1(lines []string) any {
	sum := 0

	validPull := bagPull{
		red:   12,
		green: 13,
		blue:  14,
	}

	for _, line := range lines {
		game := lineToGame(line)
		validBag := true
		for _, pull := range game.pulls {
			if pull.red > validPull.red || pull.green > validPull.green || pull.blue > validPull.blue {
				validBag = false
				break
			}
		}
		if validBag {
			sum += game.id
		}
	}

	return sum
}

func part2(lines []string) any {
	sum := 0

	for _, line := range lines {
		pulls := lineToGame(line).pulls

		minPull := bagPull{
			red:   0,
			green: 0,
			blue:  0,
		}

		for _, pull := range pulls {
			minPull.red = max(minPull.red, pull.red)
			minPull.green = max(minPull.green, pull.green)
			minPull.blue = max(minPull.blue, pull.blue)
		}

		sum += minPull.power()
	}

	return sum
}

func lineToGame(line string) game {
	split := strings.Split(line, ": ")
	gameId, err := strconv.Atoi(strings.Split(split[0], " ")[1])
	utils.Check(err)
	stringPulls := strings.Split(split[1], "; ")
	bagPulls := make([]bagPull, len(stringPulls))
	for idx, stringPull := range stringPulls {
		bagPulls[idx] = stringToPull(stringPull)
	}
	return game{
		id:    gameId,
		pulls: bagPulls,
	}
}

func stringToPull(stringPull string) bagPull {
	pull := bagPull{
		red:   0,
		green: 0,
		blue:  0,
	}
	cubes := strings.Split(stringPull, ", ")
	for _, cubeDef := range cubes {
		splitCube := strings.Split(cubeDef, " ")
		cubeCount, err := strconv.Atoi(splitCube[0])
		utils.Check(err)
		cubeColor := splitCube[1]
		switch cubeColor {
		case "red":
			pull.red = cubeCount
		case "green":
			pull.green = cubeCount
		case "blue":
			pull.blue = cubeCount
		default:
			panic(fmt.Sprintf("cube color not valid: %s", cubeColor))
		}
	}
	return pull
}

type game struct {
	id    int
	pulls []bagPull
}

type bagPull struct {
	red   int
	green int
	blue  int
}

func (b bagPull) power() int {
	return b.red * b.green * b.blue
}
