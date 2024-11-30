package day13

import (
	"advent2023/utils"
	"slices"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 13)
}

func part1(lines []string) any {
	sum := 0

	for i := 0; i < len(lines); i++ {
		pattern := [][]rune{}
		for i < len(lines) && len(lines[i]) > 0 {
			pattern = append(pattern, []rune(lines[i]))
			i++
		}
		patternReverseY := make([][]rune, len(pattern))
		for y := 0; y < len(pattern); y++ {
			patternReverseY[len(pattern)-y-1] = pattern[y]
		}
		for y := 1; y < len(pattern); y++ {
			if comparePatterns(pattern[y:], patternReverseY[len(pattern)-y:]) == 0 {
				sum += 100 * y
				break
			}
		}
		patternTransposed := make([][]rune, len(pattern[0]))
		patternReverseX := make([][]rune, len(pattern[0]))
		for x := range pattern[0] {
			patternTransposed[x] = make([]rune, len(pattern))
			patternReverseX[len(pattern[0])-x-1] = make([]rune, len(pattern))
			for y := range pattern {
				patternTransposed[x][y] = pattern[y][x]
				patternReverseX[len(pattern[0])-x-1][y] = pattern[y][x]
			}
		}
		for x := 1; x < len(pattern[0]); x++ {
			if comparePatterns(patternTransposed[x:], patternReverseX[len(pattern[0])-x:]) == 0 {
				sum += x
				break
			}
		}
	}

	return sum
}

func part2(lines []string) any {
	sum := 0

	for i := 0; i < len(lines); i++ {
		pattern := [][]rune{}
		for i < len(lines) && len(lines[i]) > 0 {
			pattern = append(pattern, []rune(lines[i]))
			i++
		}
		originalPattern := patternSol{direction: ' ', idx: 0}
		patternReverseY := make([][]rune, len(pattern))
		for y := 0; y < len(pattern); y++ {
			patternReverseY[len(pattern)-y-1] = pattern[y]
		}
		for y := 1; y < len(pattern); y++ {
			if comparePatterns(pattern[y:], patternReverseY[len(pattern)-y:]) == 0 {
				originalPattern = patternSol{direction: 'y', idx: y}
				break
			}
		}
		patternTransposed := make([][]rune, len(pattern[0]))
		patternReverseX := make([][]rune, len(pattern[0]))
		for x := range pattern[0] {
			patternTransposed[x] = make([]rune, len(pattern))
			patternReverseX[len(pattern[0])-x-1] = make([]rune, len(pattern))
			for y := range pattern {
				patternTransposed[x][y] = pattern[y][x]
				patternReverseX[len(pattern[0])-x-1][y] = pattern[y][x]
			}
		}
		for x := 1; x < len(pattern[0]); x++ {
			if comparePatterns(patternTransposed[x:], patternReverseX[len(pattern[0])-x:]) == 0 {
				originalPattern = patternSol{direction: 'x', idx: x}
				break
			}
		}
		solvedPattern := false
		for smudgeY := range pattern {
			if solvedPattern {
				break
			}
			for smudgeX := range pattern[smudgeY] {
				if pattern[smudgeY][smudgeX] == '#' {
					pattern[smudgeY][smudgeX] = '.'
				} else {
					pattern[smudgeY][smudgeX] = '#'
				}

				patternReverseY = make([][]rune, len(pattern))
				for y := 0; y < len(pattern); y++ {
					patternReverseY[len(pattern)-y-1] = pattern[y]
				}
				for y := 1; y < len(pattern); y++ {
					if comparePatterns(pattern[y:], patternReverseY[len(pattern)-y:]) == 0 && (originalPattern.direction != 'y' || originalPattern.idx != y) {
						sum += 100 * y
						solvedPattern = true
						break
					}
				}
				patternTransposed = make([][]rune, len(pattern[0]))
				patternReverseX = make([][]rune, len(pattern[0]))
				for x := range pattern[0] {
					patternTransposed[x] = make([]rune, len(pattern))
					patternReverseX[len(pattern[0])-x-1] = make([]rune, len(pattern))
					for y := range pattern {
						patternTransposed[x][y] = pattern[y][x]
						patternReverseX[len(pattern[0])-x-1][y] = pattern[y][x]
					}
				}
				for x := 1; x < len(pattern[0]); x++ {
					if comparePatterns(patternTransposed[x:], patternReverseX[len(pattern[0])-x:]) == 0 && (originalPattern.direction != 'x' || originalPattern.idx != x) {
						sum += x
						solvedPattern = true
						break
					}
				}
				if pattern[smudgeY][smudgeX] == '#' {
					pattern[smudgeY][smudgeX] = '.'
				} else {
					pattern[smudgeY][smudgeX] = '#'
				}
				if solvedPattern {
					break
				}
			}
		}
	}

	return sum
}

type patternSol struct {
	direction rune
	idx       int
}

func comparePatterns(pattern1 [][]rune, pattern2 [][]rune) int {
	for i := 0; i < min(len(pattern1), len(pattern2)); i++ {
		cmpResult := slices.Compare(pattern1[i], pattern2[i])
		if cmpResult != 0 {
			return cmpResult
		}
	}
	return 0
}
