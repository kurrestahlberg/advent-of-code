package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Test Part 1:", partOne("test.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Test 2 Part 1:", partOne("test2.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1:", partOne("input.txt"), "time spent", time.Since(start))
	//	start = time.Now()
	//	fmt.Println("Test Part 2:", partTwo("test.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", partTwo("input.txt"), "time spent", time.Since(start))
}

type Module interface {
	connectAsSource(source Module)
	connectAsDestination(destination Module)
	getSources() []Module
	getDestinations() []Module
	getId() string
	sendPulse(pulse Pulse, pulseQueue []Pulse) []Pulse
	debug() string
}

type ModuleS struct {
	id           string
	sources      []Module
	destinations []Module
}

func NewModule(name string) ModuleS {
	return ModuleS{name, []Module{}, []Module{}}
}

type Pulse struct {
	isHigh bool
	source Module
	target Module
}

var highCount int = 0
var lowCount int = 0

func NewPulse(isHigh bool, source Module, target Module) Pulse {
	// fmt.Println("New pulse", isHigh, "from", source.getId(), "to", target.getId())
	if isHigh {
		highCount += 1
	} else {
		lowCount += 1
	}

	return Pulse{isHigh, source, target}
}

func (this *ModuleS) connectAsSource(source Module) {
	this.sources = append(this.sources, source)
}

func (this *ModuleS) connectAsDestination(destination Module) {
	this.destinations = append(this.destinations, destination)
}

func (this *ModuleS) sendPulse(pulse Pulse, pulseQueue []Pulse) []Pulse {
	if !pulse.isHigh && this.getId() == "rx" {
		fmt.Println("rx at", buttonPresses+1)
	}
	for _, destination := range this.destinations {
		pulseQueue = append(pulseQueue, NewPulse(pulse.isHigh, pulse.target, destination))
	}
	return pulseQueue
}

func (this *ModuleS) getSources() []Module {
	return this.sources
}
func (this *ModuleS) getDestinations() []Module {
	return this.destinations
}
func (this *ModuleS) getId() string {
	return this.id
}

func (this *ModuleS) debug() string {
	return this.getId() + ": " + listModuleIds(this.sources) + " / " + listModuleIds(this.destinations)
}

type Broadcaster struct {
	ModuleS
}

func NewBroadcaster() *Broadcaster {
	return &Broadcaster{NewModule("broadcaster")}
}

type FlipFlop struct {
	ModuleS
	isOn bool
}

func (this *FlipFlop) sendPulse(pulse Pulse, pulseQueue []Pulse) []Pulse {
	if pulse.isHigh {
		// fmt.Println("NOT Flipping the Flop", pulse.source.getId(), this.getId())
		return pulseQueue
	}

	this.isOn = !this.isOn
	pulse.isHigh = this.isOn

	// fmt.Println("Flipping the Flop", this.isOn)

	return this.ModuleS.sendPulse(pulse, pulseQueue)
}

func (this *FlipFlop) debug() string {
	return this.ModuleS.debug() + " :: isOn = " + strconv.FormatBool(this.isOn)
}

func NewFlipFlop(name string) *FlipFlop {
	return &FlipFlop{NewModule(name), false}
}

type Conjunction struct {
	ModuleS
	memory []bool
}

func (this *Conjunction) sendPulse(pulse Pulse, pulseQueue []Pulse) []Pulse {
	for i, source := range this.sources {
		if source.getId() == pulse.source.getId() {
			this.memory[i] = pulse.isHigh
			// fmt.Println("Conjunction swapping", source.getId(), "to", pulse.isHigh)
			break
		}
	}

	output := true
	for _, mem := range this.memory {
		output = mem && output
	}

	// fmt.Println("Conjuncting", output, this.memory)

	pulse.isHigh = !output
	return this.ModuleS.sendPulse(pulse, pulseQueue)
}

func (this *Conjunction) connectAsSource(source Module) {
	this.ModuleS.connectAsSource(source)
	this.memory = append(this.memory, false)

	// fmt.Println("Conjunction memory", this.memory)
}

func (this *Conjunction) debug() string {
	return this.ModuleS.debug() + " :: memory = " + fmt.Sprint(this.memory)
}

func NewConjunction(name string) *Conjunction {
	return &Conjunction{NewModule(name), []bool{}}
}

var buttonPresses = 0

func partOne(filename string) string {
	return solve(filename, 1000)
}
func solve(filename string, rounds int) string {
	lowCount = 0
	highCount = 0
	buttonPresses = 0
	pulseQueue := []Pulse{}
	modules := make(map[string]Module)

	lines := readLines(filename)

	for _, line := range lines {
		elements := strings.Split(line, " -> ")
		switch elements[0][0] {
		case '&':
			modules[elements[0][1:]] = NewConjunction(elements[0][1:])
		case '%':
			modules[elements[0][1:]] = NewFlipFlop(elements[0][1:])
		default:
			modules[elements[0]] = NewBroadcaster()
		}
	}

	for _, line := range lines {
		elements := strings.Split(line, " -> ")
		name := elements[0]
		if name[0] == '&' || name[0] == '%' {
			name = name[1:]
		}
		destinations := strings.Split(elements[1], ", ")
		mod := modules[name]
		for _, destination := range destinations {
			// fmt.Println("Append", destination, "as destination for", mod.getId())
			if _, ok := modules[destination]; !ok {
				output := NewModule(destination)
				modules[destination] = &output
			}
			mod.connectAsDestination(modules[destination])
			modules[destination].connectAsSource(mod)
			// fmt.Println(mod.debug(), modules[destination].debug())
		}
	}

	// printModules(modules)

	b := modules["broadcaster"]
	for buttonPresses = 0; buttonPresses < rounds; buttonPresses++ {
		pulseQueue = b.sendPulse(NewPulse(false, b, b), pulseQueue)

		for len(pulseQueue) > 0 {
			pulse := pulseQueue[0]
			//fmt.Println(pulseQueue)
			pulseQueue = pulse.target.sendPulse(pulse, pulseQueue[1:])
			//printModules(modules)
		}

		// fmt.Println("Low", lowCount, "high", highCount)
		// printModules(modules)
	}
	return strconv.Itoa(lowCount * highCount)
}

func listModuleIds(modules []Module) string {
	rv := "["
	for i, mod := range modules {
		if i > 0 {
			rv += ","
		}
		rv += mod.getId()
	}
	return rv + "]"
}

func printModules(modules map[string]Module) {
	for k, v := range modules {
		fmt.Println(k, v.debug())
	}
}

func partTwo(filename string) string {
	// Dangit :D
	solve(filename, 100000000)

	return "not ready"
}
