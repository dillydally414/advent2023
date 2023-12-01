package main

import (
	"advent2023/utils"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	utils.Run(part1, part2)
}


func part1(lines []string) any {
	sum := 0

	for _, line := range lines {
		var firstDigit string
		var lastDigit string
		for _, char := range line {
			if char >= '0' && char <= '9' {
				firstDigit = string(char)
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			currChar := string(line[i])
			if currChar >= "0" && currChar <= "9" {
				lastDigit = currChar
				break
			}
		}
		value, err := strconv.Atoi(firstDigit + lastDigit)
		utils.Check(err)
		sum += value
	}

	return sum
}

func part2(lines []string) any {
	sum := 0
	numbers := make(map[string]string)
	numbers["one"] = "1"
	numbers["two"] = "2"
	numbers["three"] = "3"
	numbers["four"] = "4"
	numbers["five"] = "5"
	numbers["six"] = "6"
	numbers["seven"] = "7"
	numbers["eight"] = "8"
	numbers["nine"] = "9"

	keys := make([]string, 0, len(numbers))
	for k := range numbers {
		keys = append(keys, "(" + k + ")")
	}
	regex := regexp.MustCompile(strings.Join(keys, "|"))


	for _, line := range lines {
		var firstDigit string
		var lastDigit string
		for i, char := range line {
			if char >= '1' && char <= '9' {
				firstDigit = string(char)
				break
			}
			
			match := regex.Find([]byte(line[:i+1]))
			if match != nil {
				firstDigit = numbers[string(match)]
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			currChar := string(line[i])
			if currChar >= "0" && currChar <= "9" {
				lastDigit = currChar
				break
			} 

			match := regex.Find([]byte(line[i:]))
			if match != nil {
				lastDigit = numbers[string(match)]
				break
			}
		}
		value, err := strconv.Atoi(firstDigit + lastDigit)
		utils.Check(err)
		sum += value
	}

	return sum
}