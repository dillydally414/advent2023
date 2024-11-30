package day21

import (
	"advent2023/utils"
	"fmt"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 21)
}

func part1(lines []string) any {
	curr := []coord{}
	xMax, yMax := len(lines[0])-1, len(lines)-1
	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == 'S' {
				curr = append(curr, coord{x, y})
			}
		}
	}

	for i := 0; i < 64; i++ {
		nextMoves := make(map[coord]bool)
		for _, prev := range curr {
			for _, delta := range []coord{{x: -1, y: 0}, {x: 1, y: 0}, {x: 0, y: -1}, {x: 0, y: 1}} {
				x, y := prev.x+delta.x, prev.y+delta.y
				nextMoves[coord{x, y}] = y >= 0 && y <= yMax && x >= 0 && x <= xMax && lines[y][x] != '#'
			}
		}
		curr = []coord{}
		for k, v := range nextMoves {
			if v {
				curr = append(curr, k)
			}
		}
	}

	return len(curr)
}

func part2(lines []string) any {
	curr := make(map[coord]map[coord]bool)
	neighbors := make(map[coord][][2]coord)
	xMax, yMax := len(lines[0])-1, len(lines)-1
	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == 'S' {
				curr[coord{x, y}] = map[coord]bool{{x: 0, y: 0}: true}
			}
			if lines[y][x] != '#' {
				newNeighbors := [][2]coord{}
				for _, delta := range []coord{{x: -1, y: 0}, {x: 1, y: 0}, {x: 0, y: -1}, {x: 0, y: 1}} {
					neighborX, neighborY := (xMax+x+delta.x)%xMax, (yMax+y+delta.y)%yMax
					boardShift := coord{x: 0, y: 0}
					if neighborX < x-1 {
						boardShift.x++
					} else if neighborX > x+1 {
						boardShift.x--
					}
					if neighborY < y-1 {
						boardShift.y++
					} else if neighborY > y+1 {
						boardShift.y--
					}
					if lines[neighborY][neighborX] != '#' {
						newNeighbors = append(newNeighbors, [2]coord{{x: neighborX, y: neighborY}, boardShift})
					}
					neighbors[coord{x, y}] = newNeighbors
				}
			}
		}
	}

	for i := 0; i < 10; i++ {
		nextMoves := make(map[coord]map[coord]bool)
		for prev, boards := range curr {
			for _, neighbor := range neighbors[prev] {
				if nextMoves[neighbor[0]] == nil {
					nextMoves[neighbor[0]] = make(map[coord]bool)
				}
				for b := range boards {
					nextMoves[neighbor[0]][coord{x: neighbor[1].x + b.x, y: neighbor[1].y + b.y}] = true
				}
			}
		}
		curr = nextMoves
	}

	fmt.Println(curr)

	sum := 0

	for _, n := range curr {
		sum += len(n)
	}

	return sum
}

type coord struct {
	x int
	y int
}
