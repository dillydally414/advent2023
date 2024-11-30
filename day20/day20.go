package day20

import (
	"advent2023/utils"
	"fmt"
	"strings"
)

func GetDay() utils.Day {
	return utils.NewDay(part1, part2, 20)
}

func part1(lines []string) any {
	modules := make(map[string]module, len(lines))

	for _, line := range lines {
		sides := strings.Split(line, " -> ")
		destinations := strings.Split(sides[1], ", ")
		switch rune(sides[0][0]) {
		case 'b':
			b := broadcaster{destinations}
			modules[sides[0]] = &b
			break
		case '%':
			f := flipFlop{on: false, destinations: destinations}
			modules[sides[0][1:]] = &f
			break
		case '&':
			c := conjunction{inputPulses: make(map[string]bool), destinations: destinations}
			modules[sides[0][1:]] = &c
			break
		default:
			panic(fmt.Sprintln("Unexpected first character for input ", sides[0]))
		}
	}

	for name, m := range modules {
		for _, dest := range m.getDestinations() {
			v, ok := modules[dest].(*conjunction)
			if ok {
				v.inputPulses[name] = false
			}
		}
	}

	lowPulses := 0
	highPulses := 0

	for i := 0; i < 1000; i++ {
		pulses := []pulse{{from: "button", high: false, dest: "broadcaster"}}
		for len(pulses) > 0 {
			p := pulses[0]
			if p.high {
				highPulses++
			} else {
				lowPulses++
			}
			pulses = pulses[1:]
			if modules[p.dest] == nil {
				continue
			}
			sendPulse, pulseType := modules[p.dest].receivePulse(p.from, p.high)
			if sendPulse {
				for _, d := range modules[p.dest].getDestinations() {
					pulses = append(pulses, pulse{from: p.dest, high: pulseType, dest: d})
				}
			}
		}
	}

	fmt.Println(highPulses, lowPulses)

	return lowPulses * highPulses
}

func part2(lines []string) any {
	modules := make(map[string]module, len(lines))

	for _, line := range lines {
		sides := strings.Split(line, " -> ")
		destinations := strings.Split(sides[1], ", ")
		switch rune(sides[0][0]) {
		case 'b':
			b := broadcaster{destinations}
			modules[sides[0]] = &b
			break
		case '%':
			f := flipFlop{on: false, destinations: destinations}
			modules[sides[0][1:]] = &f
			break
		case '&':
			c := conjunction{inputPulses: make(map[string]bool), destinations: destinations}
			modules[sides[0][1:]] = &c
			break
		default:
			panic(fmt.Sprintln("Unexpected first character for input ", sides[0]))
		}
	}

	for name, m := range modules {
		for _, dest := range m.getDestinations() {
			v, ok := modules[dest].(*conjunction)
			if ok {
				v.inputPulses[name] = false
			}
		}
	}

	for buttonPresses := 0; buttonPresses >= 0; buttonPresses++ {
		pulses := []pulse{{from: "button", high: false, dest: "broadcaster"}}
		for len(pulses) > 0 {
			p := pulses[0]
			if p.dest == "rx" && !p.high {
				return buttonPresses
			}
			pulses = pulses[1:]
			if modules[p.dest] == nil {
				continue
			}
			sendPulse, pulseType := modules[p.dest].receivePulse(p.from, p.high)
			if sendPulse {
				for _, d := range modules[p.dest].getDestinations() {
					pulses = append(pulses, pulse{from: p.dest, high: pulseType, dest: d})
				}
			}
		}
	}

	return nil
}

type pulse struct {
	from string
	high bool
	dest string
}

type module interface {
	receivePulse(from string, high bool) (sendPulse bool, pulseType bool)
	getDestinations() []string
}

type broadcaster struct {
	destinations []string
}

type flipFlop struct {
	on           bool
	destinations []string
}

type conjunction struct {
	inputPulses  map[string]bool
	destinations []string
}

func (b broadcaster) getDestinations() []string {
	return b.destinations
}

func (f flipFlop) getDestinations() []string {
	return f.destinations
}

func (c conjunction) getDestinations() []string {
	return c.destinations
}

func (b *broadcaster) receivePulse(from string, high bool) (sendPulse bool, pulseType bool) {
	sendPulse = true
	pulseType = high
	return
}

func (f *flipFlop) receivePulse(from string, high bool) (sendPulse bool, pulseType bool) {
	sendPulse = !high
	if !high {
		f.on = !f.on
		pulseType = f.on
	}
	return
}

func (c *conjunction) receivePulse(from string, high bool) (sendPulse bool, pulseType bool) {
	c.inputPulses[from] = high
	sendPulse = true
	pulseType = false
	for _, prevHigh := range c.inputPulses {
		if !prevHigh {
			pulseType = true
			break
		}
	}
	return
}
