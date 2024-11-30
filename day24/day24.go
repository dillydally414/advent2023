package day24

import (
	"advent2023/utils"
	"fmt"
	"regexp"
	"strconv"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 24)
}

func part1(lines []string) any {
	hailstones := make([]hailstone, len(lines))
	slopes := make([]slope, len(lines))

	for i, line := range lines {
		hailstones[i] = parseHailstone(line)
		slopes[i] = hailstones[i].generateSlope()
	}

	min, max := 200000000000000.0, 400000000000000.0

	pathOverlap := 0

	for i, slopeA := range slopes {
		for j, slopeB := range slopes[i+1:] {
			c, parallel := slopeA.intersection(&slopeB)
			inFuture := ((hailstones[i].velocity.y >= 0 && float64(hailstones[i].position.y) <= c.y) || (hailstones[i].velocity.y < 0 && float64(hailstones[i].position.y) > c.y)) &&
				((hailstones[i+1+j].velocity.y >= 0 && float64(hailstones[i+1+j].position.y) <= c.y) || (hailstones[i+1+j].velocity.y < 0 && float64(hailstones[i+1+j].position.y) > c.y))

			if !parallel && inFuture && c.x >= min && c.x <= max && c.y >= min && c.y <= max {
				pathOverlap++
			}
		}
	}

	return pathOverlap
}

func part2(lines []string) any {
	hailstones := make([]hailstone, len(lines))
	slopes := make([]slope, len(lines))

	for i, line := range lines {
		hailstones[i] = parseHailstone(line)
		slopes[i] = hailstones[i].generateSlope()
		fmt.Println(slopes[i])
	}

	min, max := 200000000000000.0, 400000000000000.0

	pathOverlap := 0

	for i, slopeA := range slopes {
		for j, slopeB := range slopes[i+1:] {
			c, parallel := slopeA.intersection(&slopeB)
			inFuture := ((hailstones[i].velocity.y >= 0 && float64(hailstones[i].position.y) <= c.y) || (hailstones[i].velocity.y < 0 && float64(hailstones[i].position.y) > c.y)) &&
				((hailstones[i+1+j].velocity.y >= 0 && float64(hailstones[i+1+j].position.y) <= c.y) || (hailstones[i+1+j].velocity.y < 0 && float64(hailstones[i+1+j].position.y) > c.y))

			if !parallel && inFuture && c.x >= min && c.x <= max && c.y >= min && c.y <= max {
				pathOverlap++
			}
		}
	}

	return pathOverlap
}

type coord[T interface{ ~int | float64 }] struct {
	x T
	y T
	z T
}

type hailstone struct {
	position coord[int]
	velocity coord[int]
}

func parseHailstone(line string) hailstone {
	matches := regexp.MustCompile("[0-9-]+").FindAllString(line, 6)
	nums := make([]int, len(matches))
	for i, num := range matches {
		nums[i] = utils.CheckAndReturn(strconv.Atoi(num))
	}
	return hailstone{
		position: coord[int]{
			x: nums[0],
			y: nums[1],
			z: nums[2],
		},
		velocity: coord[int]{
			x: nums[3],
			y: nums[4],
			z: nums[5],
		},
	}
}

func (h *hailstone) generateSlope() slope {
	m := float64(h.velocity.y) / float64(h.velocity.x)
	b := float64(h.position.y) - m*float64(h.position.x)
	return slope{
		m,
		b,
	}
}

type slope struct {
	m float64
	b float64
}

func (s *slope) intersection(other *slope) (intersection coord[float64], parallel bool) {
	if other.m == s.m {
		parallel = true
		return
	}
	intersection.x = (s.b - other.b) / (other.m - s.m)
	intersection.y = s.m*intersection.x + s.b
	return
}
