package day23

import (
	"advent2023/utils"
	"fmt"
	"slices"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 23)
}

func part1(lines []string) any {
	paths := make(map[coord][]coord)
	isFork := make(map[coord]bool)
	isJoin := make(map[coord]bool)

	yMax := len(lines) - 1
	xMax := len(lines[0]) - 1
	var start, end coord
	for y := range lines {
		for x := range lines[y] {
			switch lines[y][x] {
			case '#':
				break
			case '>':
				paths[coord{x, y}] = []coord{{x: x + 1, y: y}}
				break
			case '<':
				paths[coord{x, y}] = []coord{{x: x - 1, y: y}}
				break
			case '^':
				paths[coord{x, y}] = []coord{{x: x, y: y - 1}}
				break
			case 'v':
				paths[coord{x, y}] = []coord{{x: x, y: y + 1}}
				break
			case '.':
				coords := []coord{}
				if y == 0 {
					start = coord{x, y}
				}
				if y == yMax {
					end = coord{x, y}
				}
				slopeInCt := 0
				slopeOutCt := 0
				if y > 0 && lines[y-1][x] != '#' {
					if lines[y-1][x] == 'v' {
						slopeInCt++
					} else {
						if lines[y-1][x] == '^' {
							slopeOutCt++
						}

						coords = append(coords, coord{x: x, y: y - 1})
					}
				}
				if y < yMax && lines[y+1][x] != '#' {
					if lines[y+1][x] == '^' {
						slopeInCt++
					} else {
						if lines[y+1][x] == 'v' {
							slopeOutCt++
						}
						coords = append(coords, coord{x: x, y: y + 1})
					}
				}
				if x > 0 && lines[y][x-1] != '#' {
					if lines[y][x-1] == '>' {
						slopeInCt++
					} else {
						if lines[y][x-1] == '<' {
							slopeOutCt++
						}
						coords = append(coords, coord{x: x - 1, y: y})
					}
				}
				if x < xMax && lines[y][x+1] != '#' {
					if lines[y][x+1] == '<' {
						slopeInCt++
					} else {
						if lines[y][x+1] == '>' {
							slopeOutCt++
						}
						coords = append(coords, coord{x: x + 1, y: y})
					}
				}
				if slopeInCt >= 1 && slopeOutCt > 1 {
					isFork[coord{x, y}] = true
				}
				if slopeInCt > 1 && slopeOutCt >= 1 {
					isJoin[coord{x, y}] = true
				}
				paths[coord{x, y}] = coords
				break
			default:
				panic(fmt.Sprintln("Unrecognized character", string(lines[y][x])))
			}
		}
	}

	importantPaths := make(map[coord][]pathTo, len(isFork)*4+len(isJoin))

	queue := []coord{start}

	for c := range isFork {
		importantPaths[c] = make([]pathTo, len(paths[c]))
		for i, out := range paths[c] {
			importantPaths[c][i] = pathTo{dest: out, length: 1}
			queue = append(queue, out)
		}
	}

	for c := range isJoin {
		if !isFork[c] {
			importantPaths[c] = make([]pathTo, len(paths[c]))
			for i, out := range paths[c] {
				importantPaths[c][i] = pathTo{dest: out, length: 1}
				queue = append(queue, out)
			}
		}
	}

	for _, curr := range queue {
		var prev coord
		lastImportant := curr
		length := 0
		for curr != end && !isFork[curr] && !isJoin[curr] {
			length++
			options := paths[curr]
			newOption := slices.IndexFunc(options, func(c coord) bool {
				return c != prev
			})
			prev = curr
			curr = options[newOption]
		}
		if importantPaths[lastImportant] == nil {
			importantPaths[lastImportant] = []pathTo{}
		}
		importantPaths[lastImportant] = append(importantPaths[lastImportant], pathTo{dest: curr, length: length})
	}

	distances := make(map[coord]int, len(importantPaths))

	queue = []coord{start}

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		neighbors := importantPaths[v]
		for _, n := range neighbors {
			queue = append(queue, n.dest)
			if distances[v]+n.length > distances[n.dest] {
				distances[n.dest] = distances[v] + n.length
			}
		}
	}

	return distances[end]
}

func part2(lines []string) any {
	paths := make(map[coord][]coord)
	isJunction := make(map[coord]bool)

	yMax := len(lines) - 1
	xMax := len(lines[0]) - 1
	var start, end coord
	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == '#' {
				continue
			}
			coords := []coord{}
			if y == 0 {
				start = coord{x, y}
			}
			if y == yMax {
				end = coord{x, y}
			}
			if y > 0 && lines[y-1][x] != '#' {
				coords = append(coords, coord{x: x, y: y - 1})

			}
			if y < yMax && lines[y+1][x] != '#' {
				coords = append(coords, coord{x: x, y: y + 1})

			}
			if x > 0 && lines[y][x-1] != '#' {
				coords = append(coords, coord{x: x - 1, y: y})

			}
			if x < xMax && lines[y][x+1] != '#' {
				coords = append(coords, coord{x: x + 1, y: y})

			}
			if len(coords) > 2 {
				isJunction[coord{x, y}] = true
			}
			paths[coord{x, y}] = coords
		}
	}

	importantPaths := make(map[coord][]pathTo, len(isJunction)*5)

	queue := []coord{start}

	for c := range isJunction {
		importantPaths[c] = make([]pathTo, len(paths[c]))
		for i, out := range paths[c] {
			importantPaths[c][i] = pathTo{dest: out, length: 1}
			queue = append(queue, out)
		}
	}

	for _, curr := range queue {
		var prev coord
		lastImportant := curr
		length := 0
		for curr != end && !(length > 0 && curr == start) && !isJunction[curr] {
			length++
			options := paths[curr]
			newOption := slices.IndexFunc(options, func(c coord) bool {
				return c != prev
			})
			prev = curr
			curr = options[newOption]
		}
		if importantPaths[lastImportant] == nil {
			importantPaths[lastImportant] = []pathTo{}
		}
		importantPaths[lastImportant] = append(importantPaths[lastImportant], pathTo{dest: curr, length: length})
	}

	distances := make(map[coord]int, len(importantPaths))

	queue = []coord{start}

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		neighbors := importantPaths[v]
		for _, n := range neighbors {
			queue = append(queue, n.dest)
			if distances[v]+n.length > distances[n.dest] {
				distances[n.dest] = distances[v] + n.length
			}
		}
	}

	return distances[end]
}

type coord struct {
	x int
	y int
}

type pathTo struct {
	dest   coord
	length int
}
