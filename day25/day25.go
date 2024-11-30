package day25

import (
	"advent2023/utils"
	"math/rand"
	"regexp"
	"slices"
	"strings"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 25)
}

func part1(lines []string) any {
	components := make(map[string][]string, len(lines))
	edges := [][2]string{}
	for _, line := range lines {
		split := regexp.MustCompile(":? ").Split(line, len(line))
		if components[split[0]] == nil {
			components[split[0]] = split[1:]
		} else {
			components[split[0]] = append(components[split[0]], split[1:]...)
		}
		for _, c := range split[1:] {
			edge := [2]string{split[0], c}
			if strings.Compare(split[0], c) > 0 {
				edge[0] = c
				edge[1] = split[0]
			}
			edges = append(edges, edge)
			if components[c] == nil {
				components[c] = []string{split[0]}
			} else {
				components[c] = append(components[c], split[0])
			}
		}
	}

	for {
		representatives := make(map[string]string)
		newEdges := make(map[string][][2]string, len(components))

		for c, neighbors := range components {
			newEdges[c] = make([][2]string, len(neighbors))
			for i, neighbor := range neighbors {
				newEdges[c][i] = [2]string{c, neighbor}
			}
		}

		rand.Shuffle(len(edges), func(i, j int) {
			edges[i], edges[j] = edges[j], edges[i]
		})

		for len(newEdges) > 2 {
			var e [2]string
			for _, v := range newEdges {
				for _, newE := range v {
					e = newE
					break
				}
				break
			}
			rep1, rep2 := retrieveRep(representatives, e[0]), retrieveRep(representatives, e[1])
			if rep1 == rep2 {
				continue
			}
			representatives[rep1] = rep2
			newRep2Edges := [][2]string{}
			for _, n := range newEdges[rep1] {
				if !slices.Contains(newEdges[rep2], n) && (retrieveRep(representatives, n[0]) != rep2 || retrieveRep(representatives, n[1]) != rep2) {
					newRep2Edges = append(newRep2Edges, n)
				}
			}
			for _, n := range newEdges[rep2] {
				if retrieveRep(representatives, n[0]) != rep2 || retrieveRep(representatives, n[1]) != rep2 {
					newRep2Edges = append(newRep2Edges, n)
				}
			}
			newEdges[rep2] = newRep2Edges
			delete(newEdges, rep1)
		}

		for _, v := range newEdges {
			if len(v) == 3 {
				cts := make(map[string]int, 2)
				for k := range representatives {
					cts[retrieveRep(representatives, k)]++
				}
				product := 1
				for _, num := range cts {
					product *= num + 1
				}
				return product
			}
			break
		}
	}
}

func part2(lines []string) any {
	return nil
}

func retrieveRep(representatives map[string]string, component string) string {
	curr := component
	for {
		if _, ok := representatives[curr]; ok {
			curr = representatives[curr]
		} else {
			return curr
		}
	}
}
