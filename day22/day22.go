package day22

import (
	"advent2023/utils"
	"slices"
	"strconv"
	"strings"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 22)
}

func part1(lines []string) any {
	bricks := make([]brick, len(lines))

	for i, line := range lines {
		sides := strings.Split(line, "~")
		left, right := strings.Split(sides[0], ","), strings.Split(sides[1], ",")
		bricks[i] = brick{
			x: [2]int{utils.CheckAndReturn(strconv.Atoi(left[0])), utils.CheckAndReturn(strconv.Atoi(right[0]))},
			y: [2]int{utils.CheckAndReturn(strconv.Atoi(left[1])), utils.CheckAndReturn(strconv.Atoi(right[1]))},
			z: [2]int{utils.CheckAndReturn(strconv.Atoi(left[2])), utils.CheckAndReturn(strconv.Atoi(right[2]))},
		}
	}

	slices.SortFunc(bricks, func(i, j brick) int {
		return i.z[0] - j.z[0]
	})

	baseline := 1
	for i, b := range bricks {
		b.z[1] -= b.z[0] - baseline
		b.z[0] = baseline
		for newBottom := baseline; newBottom > 1; newBottom-- {
			settled := false
			for _, prevBrick := range bricks[0:i] {
				settled = settled || prevBrick.supports(b)
			}
			if settled {
				break
			}
			b.z[1]--
			b.z[0]--
		}
		baseline = max(baseline, b.z[1]+1)
		bricks[i] = b
	}

	slices.SortFunc(bricks, func(i, j brick) int {
		return i.z[0] - j.z[0]
	})

	supportedBy := make(map[brick][]brick, len(bricks))
	supportsOther := make(map[brick][]brick, len(bricks))

	for _, b := range bricks {
		supportedBy[b] = []brick{}
		supportsOther[b] = []brick{}
	}

	for _, b := range bricks {
		for _, other := range bricks {
			if other.supports(b) {
				supportedBy[b] = append(supportedBy[b], other)
			} else if b.supports(other) {
				supportsOther[b] = append(supportsOther[b], other)
			}
		}
	}

	ct := 0

	for _, b := range bricks {
		canRemove := true
		for _, reliant := range supportsOther[b] {
			if len(supportedBy[reliant]) == 1 {
				canRemove = false
			}
		}
		if canRemove {
			ct++
		}
	}

	return ct
}

func part2(lines []string) any {
	bricks := make([]brick, len(lines))

	for i, line := range lines {
		sides := strings.Split(line, "~")
		left, right := strings.Split(sides[0], ","), strings.Split(sides[1], ",")
		bricks[i] = brick{
			x: [2]int{utils.CheckAndReturn(strconv.Atoi(left[0])), utils.CheckAndReturn(strconv.Atoi(right[0]))},
			y: [2]int{utils.CheckAndReturn(strconv.Atoi(left[1])), utils.CheckAndReturn(strconv.Atoi(right[1]))},
			z: [2]int{utils.CheckAndReturn(strconv.Atoi(left[2])), utils.CheckAndReturn(strconv.Atoi(right[2]))},
		}
	}

	slices.SortFunc(bricks, func(i, j brick) int {
		return i.z[0] - j.z[0]
	})

	baseline := 1
	for i, b := range bricks {
		b.z[1] -= b.z[0] - baseline
		b.z[0] = baseline
		for newBottom := baseline; newBottom > 1; newBottom-- {
			settled := false
			for _, prevBrick := range bricks[0:i] {
				settled = settled || prevBrick.supports(b)
			}
			if settled {
				break
			}
			b.z[1]--
			b.z[0]--
		}
		baseline = max(baseline, b.z[1]+1)
		bricks[i] = b
	}

	slices.SortFunc(bricks, func(i, j brick) int {
		return i.z[0] - j.z[0]
	})

	supportedBy := make(map[brick][]brick, len(bricks))
	supportsOther := make(map[brick][]brick, len(bricks))

	for _, b := range bricks {
		supportedBy[b] = []brick{}
		supportsOther[b] = []brick{}
	}

	for _, b := range bricks {
		for _, other := range bricks {
			if other.supports(b) {
				supportedBy[b] = append(supportedBy[b], other)
			} else if b.supports(other) {
				supportsOther[b] = append(supportsOther[b], other)
			}
		}
	}

	ct := 0

	for i, disintegrated := range bricks {
		wouldFall := make(map[brick]bool, len(bricks)-i)
		fallQueue := []brick{disintegrated}
		wouldFall[disintegrated] = true
		for len(fallQueue) > 0 {
			b := fallQueue[0]
			fallQueue = fallQueue[1:]
			wouldBrickFall := true
			for _, underneath := range supportedBy[b] {
				wouldBrickFall = wouldBrickFall && wouldFall[underneath]
			}
			wouldFall[b] = wouldBrickFall || b == disintegrated
			if wouldFall[b] {
				fallQueue = append(fallQueue, supportsOther[b]...)
			}
		}
		wouldFall[disintegrated] = false // it's being disintegrated not falling
		for _, v := range wouldFall {
			if v {
				ct++
			}
		}
	}

	return ct
}

type brick struct {
	x [2]int
	y [2]int
	z [2]int
}

func (b *brick) supports(other brick) bool {
	return b.z[1] == other.z[0]-1 && (b.x[0] <= other.x[1] && b.x[1] >= other.x[0]) && (b.y[0] <= other.y[1] && b.y[1] >= other.y[0])
}
