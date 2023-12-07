package day7

import (
	"advent2023/utils"
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 7)
}

func part1(lines []string) any {
	hands := make([]hand, len(lines))

	for i, line := range lines {
		handsSplit := strings.Split(line, " ")
		strCards, bid := strings.Split(handsSplit[0], ""), utils.CheckAndReturn(strconv.Atoi(handsSplit[1]))
		cards := make([]card, len(strCards))
		for c, str := range strCards {
			cards[c] = card{str}
		}
		hands[i] = hand{
			bid,
			cards,
		}
	}

	slices.SortFunc(hands, compareHands)

	winnings := 0

	for i, h := range hands {
		winnings += h.bid * (i + 1)
	}

	return winnings
}

func part2(lines []string) any {
	hands := make([]hand, len(lines))

	for i, line := range lines {
		handsSplit := strings.Split(line, " ")
		strCards, bid := strings.Split(handsSplit[0], ""), utils.CheckAndReturn(strconv.Atoi(handsSplit[1]))
		cards := make([]card, len(strCards))
		for c, str := range strCards {
			cards[c] = card{str}
		}
		hands[i] = hand{
			bid,
			cards,
		}
	}

	slices.SortFunc(hands, compareHands2)

	winnings := 0

	for i, h := range hands {
		winnings += h.bid * (i + 1)
	}

	return winnings
}

type card struct {
	string
}

type hand struct {
	bid   int
	cards []card
}

var HAND_ORDER = [...]string{"five", "four", "full house", "three", "two pair", "pair", "high card"}
var CARD_ORDER = [...]string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

func compareCards(c1 card, c2 card) int {
	return cmp.Compare(slices.Index(CARD_ORDER[:], c1.string), slices.Index(CARD_ORDER[:], c2.string))
}

func (h hand) handType(jWild bool) string {
	frequencies := make(map[string]int)
	for _, c := range h.cards {
		frequencies[c.string] += 1
	}
	maxCount, secondMaxCount := 0, 0
	for k, v := range frequencies {
		if k == "J" {
			continue
		}
		if v > maxCount {
			secondMaxCount = maxCount
			maxCount = v
		} else if v > secondMaxCount {
			secondMaxCount = v
		}
	}
	maxCount += frequencies["J"]
	switch maxCount {
	case 1:
		return "high card"
	case 2:
		if secondMaxCount == 2 {
			return "two pair"
		} else {
			return "pair"
		}
	case 3:
		if secondMaxCount == 2 {
			return "full house"
		} else {
			return "three"
		}
	case 4:
		return "four"
	case 5:
		return "five"
	default:
		panic(fmt.Sprintf("Count invalid: %d", maxCount))
	}
}

func compareHands(h1 hand, h2 hand) int {
	typeComparison := cmp.Compare(slices.Index(HAND_ORDER[:], h1.handType(false)), slices.Index(HAND_ORDER[:], h2.handType(false))) * -1
	if typeComparison != 0 {
		return typeComparison
	} else {
		for i := range h1.cards {
			cardComparison := compareCards(h1.cards[i], h2.cards[i]) * -1
			if cardComparison != 0 {
				return cardComparison
			}
		}
	}
	return 0
}

func compareHands2(h1 hand, h2 hand) int {
	typeComparison := cmp.Compare(slices.Index(HAND_ORDER[:], h1.handType(true)), slices.Index(HAND_ORDER[:], h2.handType(true))) * -1
	if typeComparison != 0 {
		return typeComparison
	} else {
		for i := range h1.cards {
			cardComparison := compareCards(h1.cards[i], h2.cards[i]) * -1
			if cardComparison != 0 {
				if h1.cards[i].string == "J" {
					return -1
				} else if h2.cards[i].string == "J" {
					return 1
				} else {
					return cardComparison
				}
			}
		}
	}
	return 0
}
