package day17

import (
	"advent2023/utils"
	"math"
	"strconv"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 17)
}

var neighborDelta = [4]coord{{x: 1, y: 0}, {x: -1, y: 0}, {x: 0, y: -1}, {x: 0, y: 1}}

func part1(lines []string) any {
	visited := make(map[blockMove]bool, len(lines)*len(lines[0]))
	distances := make(map[blockMove]int, len(lines)*len(lines[0]))
	fscore := make(map[blockMove]int, len(lines)*len(lines[0]))
	weights := make(map[coord]int, len(lines)*len(lines[0]))
	prev := make(map[blockMove]blockMove, len(lines)*len(lines[0]))
	zeroPrev := coord{x: -1, y: -1}
	source := blockMove{
		loc: coord{x: 0, y: 0},
	}
	dest := coord{x: len(lines[0]) - 1, y: len(lines) - 1}

	for y, line := range lines {
		for x, r := range line {
			c := coord{x, y}
			weights[c] = utils.CheckAndReturn(strconv.Atoi(string(r)))
		}
	}
	destMoves := []blockMove{{loc: dest, right: 1}, {loc: dest, right: 2}, {loc: dest, right: 3}, {loc: dest, down: 1}, {loc: dest, down: 2}, {loc: dest, down: 3}}
	distances[source] = 0
	fscore[source] = 0

	for {
		u := blockMove{loc: zeroPrev}
		for k, v := range fscore {
			if !visited[k] && (u.loc == zeroPrev || v < fscore[u]) {
				u = k
			}
		}
		visited[u] = true
		delete(fscore, u)
		if u.loc == zeroPrev || u.loc == dest {
			break
		}

		for _, v := range u.neighbors(len(lines[0]), len(lines)) {
			if visited[v] || v.loc.x < 0 || v.loc.x > dest.x || v.loc.y < 0 || v.loc.y > dest.y {
				continue
			}
			newDist := distances[u] + weights[v.loc]
			oldDist, existed := distances[v]
			if !existed || newDist < oldDist {
				distances[v] = newDist
				prev[v] = u
				fscore[v] = newDist + dest.manhattanDist(v.loc)
			}
		}
	}

	minDist := 1000000
	for _, b := range destMoves {
		dist, existed := distances[b]
		if !existed {
			continue
		}
		minDist = min(dist, minDist)
	}

	return minDist
}

func part2(lines []string) any {
	visited := make(map[blockMove]bool, len(lines)*len(lines[0]))
	distances := make(map[blockMove]int, len(lines)*len(lines[0]))
	fscore := make(map[blockMove]int, len(lines)*len(lines[0]))
	weights := make(map[coord]int, len(lines)*len(lines[0]))
	prev := make(map[blockMove]blockMove, len(lines)*len(lines[0]))
	zeroPrev := coord{x: -1, y: -1}
	sourceDown := blockMove{
		loc:  coord{x: 0, y: 0},
		down: 1,
	}
	sourceRight := blockMove{
		loc:   coord{x: 0, y: 0},
		right: 1,
	}
	dest := coord{x: len(lines[0]) - 1, y: len(lines) - 1}

	for y, line := range lines {
		for x, r := range line {
			c := coord{x, y}
			weights[c] = utils.CheckAndReturn(strconv.Atoi(string(r)))
		}
	}

	destMoves := []blockMove{}
	for i := 4; i <= 10; i++ {
		destMoves = append(destMoves, blockMove{loc: dest, right: i}, blockMove{loc: dest, down: i})
	}
	distances[sourceDown] = 0
	fscore[sourceDown] = 0
	distances[sourceRight] = 0
	fscore[sourceRight] = 0

	for {
		u := blockMove{loc: zeroPrev}
		for k, v := range fscore {
			if !visited[k] && (u.loc == zeroPrev || v < fscore[u]) {
				u = k
			}
		}
		visited[u] = true
		delete(fscore, u)
		if u.loc == zeroPrev || u.loc == dest {
			break
		}

		for _, v := range u.neighbors2(len(lines[0]), len(lines)) {
			if visited[v] || v.loc.x < 0 || v.loc.x > dest.x || v.loc.y < 0 || v.loc.y > dest.y {
				continue
			}
			newDist := distances[u] + weights[v.loc]
			oldDist, existed := distances[v]
			if !existed || newDist < oldDist {
				distances[v] = newDist
				prev[v] = u
				fscore[v] = newDist + dest.manhattanDist(v.loc)
			}
		}
	}

	minDist := 1000000
	for _, b := range destMoves {
		dist, existed := distances[b]
		if !existed {
			continue
		}
		minDist = min(dist, minDist)
	}

	return minDist
}

type blockMove struct {
	loc                   coord
	up, down, left, right int
}

type coord struct {
	x int
	y int
}

func (b *blockMove) neighbors(xMax int, yMax int) []blockMove {
	neighbors := []blockMove{}
	adjustments := []blockMove{}
	inline := b.inLine()
	if !inline || b.right+b.left > 0 {
		if b.up == 0 {
			adjustments = append(adjustments, blockMove{loc: coord{y: 1}, down: b.down + 1})
		}
		if b.down == 0 {
			adjustments = append(adjustments, blockMove{loc: coord{y: -1}, up: b.up + 1})
		}
	}
	if !inline || b.up+b.down > 0 {
		if b.left == 0 {
			adjustments = append(adjustments, blockMove{loc: coord{x: 1}, right: b.right + 1})
		}
		if b.right == 0 {
			adjustments = append(adjustments, blockMove{loc: coord{x: -1}, left: b.left + 1})
		}
	}
	for _, adjustment := range adjustments {
		if b.loc.x+adjustment.loc.x >= 0 && b.loc.y+adjustment.loc.y >= 0 && b.loc.x+adjustment.loc.x < xMax && b.loc.y+adjustment.loc.y < yMax {
			adjustment.loc = coord{x: b.loc.x + adjustment.loc.x, y: b.loc.y + adjustment.loc.y}
			neighbors = append(neighbors, adjustment)
		}
	}
	return neighbors
}

func (b *blockMove) inLine() bool {
	return b.up >= 3 || b.down >= 3 || b.left >= 3 || b.right >= 3
}

func (c *coord) manhattanDist(other coord) int {
	return int(math.Abs(float64(c.x-other.y)) + math.Abs(float64(c.y-other.y)))
}

func (b *blockMove) neighbors2(xMax int, yMax int) []blockMove {
	neighbors := []blockMove{}
	adjustments := []blockMove{}
	mustTurn := b.up >= 10 || b.down >= 10 || b.left >= 10 || b.right >= 10
	cantTurn := b.up < 4 && b.down < 4 && b.left < 4 && b.right < 4
	if mustTurn {
		if b.up+b.down == 0 {
			adjustments = append(adjustments, blockMove{loc: coord{y: 1}, down: b.down + 1}, blockMove{loc: coord{y: -1}, up: b.up + 1})
		}
		if b.left+b.right == 0 {
			adjustments = append(adjustments, blockMove{loc: coord{x: 1}, right: b.right + 1}, blockMove{loc: coord{x: -1}, left: b.left + 1})
		}
	} else if cantTurn {
		if b.up != 0 {
			adjustments = append(adjustments, blockMove{loc: coord{y: -1}, up: b.up + 1})
		}
		if b.down != 0 {
			adjustments = append(adjustments, blockMove{loc: coord{y: 1}, down: b.down + 1})
		}
		if b.left != 0 {
			adjustments = append(adjustments, blockMove{loc: coord{x: -1}, left: b.left + 1})
		}
		if b.right != 0 {
			adjustments = append(adjustments, blockMove{loc: coord{x: 1}, right: b.right + 1})
		}
	} else {
		if b.up == 0 {
			adjustments = append(adjustments, blockMove{loc: coord{y: 1}, down: b.down + 1})
		}
		if b.down == 0 {
			adjustments = append(adjustments, blockMove{loc: coord{y: -1}, up: b.up + 1})
		}
		if b.left == 0 {
			adjustments = append(adjustments, blockMove{loc: coord{x: 1}, right: b.right + 1})
		}
		if b.right == 0 {
			adjustments = append(adjustments, blockMove{loc: coord{x: -1}, left: b.left + 1})
		}
	}
	for _, adjustment := range adjustments {
		if b.loc.x+adjustment.loc.x >= 0 && b.loc.y+adjustment.loc.y >= 0 && b.loc.x+adjustment.loc.x < xMax && b.loc.y+adjustment.loc.y < yMax {
			adjustment.loc = coord{x: b.loc.x + adjustment.loc.x, y: b.loc.y + adjustment.loc.y}
			neighbors = append(neighbors, adjustment)
		}
	}
	return neighbors
}
