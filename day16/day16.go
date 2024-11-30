package day16

import (
	"advent2023/utils"
	"fmt"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 16)
}

func part1(lines []string) any {
	xLen := len(lines[0])
	yLen := len(lines)
	energized := make(map[coord]bool)
	visitedDirections := make(map[movement]bool)
	currMovements := []movement{{location: coord{x: 0, y: 0}, direction: coord{x: 1, y: 0}}}

	for len(currMovements) > 0 {
		curr := currMovements[0]
		if visitedDirections[curr] || curr.location.x < 0 || curr.location.x >= xLen || curr.location.y < 0 || curr.location.y >= yLen {
			currMovements = currMovements[1:]
			continue
		}
		tile := rune(lines[curr.location.y][curr.location.x])
		visitedDirections[curr] = true
		energized[curr.location] = true
		switch tile {
		case '\\':
			temp := curr.direction.x
			curr.direction.x = curr.direction.y
			curr.direction.y = temp
			break
		case '/':
			temp := curr.direction.x
			curr.direction.x = -curr.direction.y
			curr.direction.y = -temp
			break
		case '|':
			if curr.direction.y == 0 {
				curr.direction.x = 0
				curr.direction.y = 1
				currMovements = append(currMovements, movement{
					location:  curr.location,
					direction: coord{x: 0, y: -1},
				})
			}
			break
		case '-':
			if curr.direction.x == 0 {
				curr.direction.y = 0
				curr.direction.x = 1
				currMovements = append(currMovements, movement{
					location:  curr.location,
					direction: coord{x: -1, y: 0},
				})
			}
			break
		case '.':
			break
		default:
			panic(fmt.Sprintln("Unrecognized symbol: ", tile))
		}
		curr.location.x += curr.direction.x
		curr.location.y += curr.direction.y
		currMovements[0] = curr
	}

	return len(energized)
}

func part2(lines []string) any {
	xLen := len(lines[0])
	yLen := len(lines)
	maxEnergized := 0

	startingMovements := make([]movement, 2*xLen+2*yLen)

	for x := 0; x < xLen; x++ {
		startingMovements[x] = movement{location: coord{
			x: x,
			y: 0,
		}, direction: coord{
			x: 0,
			y: 1,
		}}
		startingMovements[x+xLen] = movement{location: coord{
			x: x,
			y: yLen - 1,
		}, direction: coord{
			x: 0,
			y: -1,
		}}
	}

	for y := 0; y < yLen; y++ {
		startingMovements[2*xLen+y] = movement{location: coord{
			x: 0,
			y: y,
		}, direction: coord{
			x: 1,
			y: 0,
		}}
		startingMovements[2*xLen+y+yLen] = movement{location: coord{
			x: xLen - 1,
			y: y,
		}, direction: coord{
			x: -1,
			y: 0,
		}}
	}

	for _, startingMovement := range startingMovements {

		energized := make(map[coord]bool)
		visitedDirections := make(map[movement]bool)
		currMovements := []movement{startingMovement}

		for len(currMovements) > 0 {
			curr := currMovements[0]
			if visitedDirections[curr] || curr.location.x < 0 || curr.location.x >= xLen || curr.location.y < 0 || curr.location.y >= yLen {
				currMovements = currMovements[1:]
				continue
			}
			tile := rune(lines[curr.location.y][curr.location.x])
			visitedDirections[curr] = true
			energized[curr.location] = true
			switch tile {
			case '\\':
				temp := curr.direction.x
				curr.direction.x = curr.direction.y
				curr.direction.y = temp
				break
			case '/':
				temp := curr.direction.x
				curr.direction.x = -curr.direction.y
				curr.direction.y = -temp
				break
			case '|':
				if curr.direction.y == 0 {
					curr.direction.x = 0
					curr.direction.y = 1
					currMovements = append(currMovements, movement{
						location:  curr.location,
						direction: coord{x: 0, y: -1},
					})
				}
				break
			case '-':
				if curr.direction.x == 0 {
					curr.direction.y = 0
					curr.direction.x = 1
					currMovements = append(currMovements, movement{
						location:  curr.location,
						direction: coord{x: -1, y: 0},
					})
				}
				break
			case '.':
				break
			default:
				panic(fmt.Sprintln("Unrecognized symbol: ", tile))
			}
			curr.location.x += curr.direction.x
			curr.location.y += curr.direction.y
			currMovements[0] = curr
		}
		maxEnergized = max(maxEnergized, len(energized))
	}

	return maxEnergized
}

type coord struct {
	x int
	y int
}

type movement struct {
	location  coord
	direction coord
}
