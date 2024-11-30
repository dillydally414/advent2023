package day19

import (
	"advent2023/utils"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 19)
}

func part1(lines []string) any {
	workflows := make(map[string][]rule)
	parts := []part{}
	readingParts := false

	for _, line := range lines {
		if len(line) == 0 {
			readingParts = true
			continue
		}
		if readingParts {
			split := regexp.MustCompile("[0-9]+").FindAllString(line, len(line))
			parts = append(parts, part{
				x: utils.CheckAndReturn(strconv.Atoi(split[0])),
				m: utils.CheckAndReturn(strconv.Atoi(split[1])),
				a: utils.CheckAndReturn(strconv.Atoi(split[2])),
				s: utils.CheckAndReturn(strconv.Atoi(split[3])),
			})
		} else {
			name, rules := parseWorkflow(line)
			workflows[name] = rules
		}
	}

	sum := 0

	for _, p := range parts {

		if p.partAccepted(workflows) {
			sum += p.sum()
		}

	}

	return sum
}

func part2(lines []string) any {
	workflows := make(map[string][]rule)
	breakpoints := make(map[rune][]int, 4)
	breakpoints['x'] = []int{0}
	breakpoints['m'] = []int{0}
	breakpoints['a'] = []int{0}
	breakpoints['s'] = []int{0}

	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		name, rules := parseWorkflow(line)
		workflows[name] = rules
		x, m, a, s := getBreakpoints(line)
		breakpoints['x'] = append(breakpoints['x'], x...)
		breakpoints['m'] = append(breakpoints['m'], m...)
		breakpoints['a'] = append(breakpoints['a'], a...)
		breakpoints['s'] = append(breakpoints['s'], s...)
	}

	breakpoints['x'] = append(breakpoints['x'], 4000)
	breakpoints['m'] = append(breakpoints['m'], 4000)
	breakpoints['a'] = append(breakpoints['a'], 4000)
	breakpoints['s'] = append(breakpoints['s'], 4000)

	slices.Sort(breakpoints['x'])
	slices.Sort(breakpoints['m'])
	slices.Sort(breakpoints['a'])
	slices.Sort(breakpoints['s'])

	fmt.Println(len(breakpoints['x']), len(breakpoints['m']), len(breakpoints['a']), len(breakpoints['s']), len(breakpoints['x'])*len(breakpoints['m'])*len(breakpoints['a'])*len(breakpoints['s']))

	acceptable := 0

	for xIdx, x := range breakpoints['x'][1:] {
		for mIdx, m := range breakpoints['m'][1:] {
			for aIdx, a := range breakpoints['a'][1:] {
				for sIdx, s := range breakpoints['s'][1:] {
					p := part{x, m, a, s}
					if p.partAccepted(workflows) {
						acceptable += (x - breakpoints['x'][xIdx]) * (m - breakpoints['m'][mIdx]) * (a - breakpoints['a'][aIdx]) * (s - breakpoints['s'][sIdx])
					}
				}
			}
		}
	}

	return acceptable
}

type part struct {
	x int
	m int
	a int
	s int
}

func (p *part) sum() int {
	return p.x + p.m + p.a + p.s
}

type rule struct {
	apply func(p *part) string
}

func parseWorkflow(input string) (name string, rules []rule) {
	split := regexp.MustCompile("{|}").Split(input, len(input))
	name = split[0]
	rules = []rule{}
	for _, r := range strings.Split(split[1], ",") {
		ruleSplit := regexp.MustCompile("<|>|:").Split(r, len(r))
		if len(ruleSplit) == 1 {
			rules = append(rules, rule{
				apply: func(p *part) string { return ruleSplit[0] },
			})
		} else {
			ruleStr := strings.Clone(r)
			category := ruleSplit[0]
			num := utils.CheckAndReturn(strconv.Atoi(ruleSplit[1]))
			dest := ruleSplit[2]
			rules = append(rules, rule{
				apply: func(p *part) string {
					var partVal int
					switch category {
					case "x":
						partVal = p.x
						break
					case "m":
						partVal = p.m
						break
					case "a":
						partVal = p.a
						break
					case "s":
						partVal = p.s
						break
					default:
						panic(fmt.Sprintln("Unsupported category: ", category))
					}
					var matches bool
					if strings.Contains(ruleStr, "<") {
						matches = partVal < num
					} else {
						matches = partVal > num
					}
					if matches {
						return dest
					} else {
						return ""
					}
				},
			})
		}
	}
	return
}

func (p *part) partAccepted(workflows map[string][]rule) bool {
	currWorkflow := "in"
	for {
		rules := workflows[currWorkflow]
		foundNext := false
		for _, r := range rules {
			next := r.apply(p)
			switch next {
			case "A":
				return true
			case "R":
				return false
			case "":
				break
			default:
				currWorkflow = next
				foundNext = true
				break
			}
			if foundNext {
				break
			}
		}
	}
}

func getBreakpoints(input string) (x []int, m []int, a []int, s []int) {
	split := regexp.MustCompile("{|}").Split(input, len(input))
	x = []int{}
	m = []int{}
	a = []int{}
	s = []int{}
	for _, r := range strings.Split(split[1], ",") {
		ruleSplit := regexp.MustCompile("<|>|:").Split(r, len(r))
		if len(ruleSplit) == 3 {
			category := ruleSplit[0]
			num := utils.CheckAndReturn(strconv.Atoi(ruleSplit[1]))
			if strings.Contains(r, "<") {
				num--
			}
			switch category {
			case "x":
				x = append(x, num)
				break
			case "m":
				m = append(m, num)
				break
			case "a":
				a = append(a, num)
				break
			case "s":
				s = append(s, num)
				break
			default:
				panic(fmt.Sprintln("Unsupported category: ", category))
			}
		}
	}
	return
}
