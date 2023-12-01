package utils

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Part func(lines []string) (output any)

// helper function to panic on error
func Check(e error) {
    if e != nil {
        panic(e)
    }
}

func determineFileName(actual bool, part2 bool) (filename string) {
	if actual  { 
		return "actual_input.txt"
	} else { 
		if _, err := os.Stat("sample_input_2.txt"); part2 && !os.IsNotExist(err) {
			return "sample_input_2.txt"
		} else {
			return "sample_input.txt"
		}
	}
}

// abstract runner that passes line-split input into part1 and part2 handlers, with inputs
// determined by command-line flags
func Run(part1 Part, part2 Part) {
	actualPtr := flag.Bool("actual", false, "run on actual input")
	part2Ptr := flag.Bool("part2", false, "run part 2")
	flag.Parse()

	f, err := os.Open(determineFileName(*actualPtr, *part2Ptr))
	Check(err)

	scanner := bufio.NewScanner(f)

	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	
	var output any

	if *part2Ptr {
		output = part2(lines)
	} else {
		output = part1(lines)
	}

	fmt.Println(output)
}