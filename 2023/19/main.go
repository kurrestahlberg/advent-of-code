package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Test Part 1:", partOne("test.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1:", partOne("input.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Test Part 2:", partTwo("test.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", partTwo("input.txt"), "time spent", time.Since(start))
}

type Part map[string]int

var partRegex = regexp.MustCompile(`([xmas])=([0-9]+)`)

func newPart(partDesc string) Part {
	part := Part{}
	match := partRegex.FindAllStringSubmatch(partDesc, -1)
	for _, val := range match {
		part[val[1]], _ = strconv.Atoi(val[2])
	}
	return part
}

type Rule struct {
	category string
	op       string
	value    int
	target   string
}

func (r Rule) try(part Part) bool {
	val := part[r.category]

	if r.op == ">" {
		if val > r.value {
			return true
		}
	} else {
		if val < r.value {
			return true
		}
	}

	return false
}

var ruleRegex = regexp.MustCompile(`([xmas])([<>])([0-9]+):([a-z]+|[AR])`)

func newRule(ruleDescription string) Rule {
	match := ruleRegex.FindStringSubmatch(ruleDescription)
	if len(match) > 0 {
		val, _ := strconv.Atoi(match[3])
		return Rule{match[1], match[2], val, match[4]}
	}

	return Rule{"x", ">", -1, ruleDescription}
}

type Workflow struct {
	name  string
	rules []Rule
}

var workflowRegex = regexp.MustCompile(`([a-z]+)\{([^}]+)\}`)

func newWorkflow(workflowDesc string) Workflow {
	match := workflowRegex.FindStringSubmatch(workflowDesc)
	rules := strings.Split(match[2], ",")
	workflow := Workflow{match[1], make([]Rule, len(rules))}
	for i, rule := range rules {
		workflow.rules[i] = newRule(rule)
	}
	return workflow
}

func partOne(filename string) string {

	workflows, parts := parse(filename)

	accepted := []Part{}

	for _, part := range parts {
		if sortPart(workflows, part) {
			accepted = append(accepted, part)
		}
	}

	return strconv.Itoa(addValues(accepted))
}

func parse(filename string) (map[string]Workflow, []Part) {
	lines := readLines(filename)

	workflows := make(map[string]Workflow)
	parts := []Part{}

	parsingWorkflows := true
	for _, line := range lines {
		if parsingWorkflows {
			if len(line) == 0 {
				parsingWorkflows = false
			} else {
				wf := newWorkflow(line)
				workflows[wf.name] = wf
			}
		} else {
			parts = append(parts, newPart(line))
		}
	}
	return workflows, parts
}

func addValues(parts []Part) int {
	total := 0
	for _, part := range parts {
		total += part["x"] + part["m"] + part["a"] + part["s"]
	}
	return total
}

func sortPart(workflows map[string]Workflow, part Part) bool {
	currentWorkflow := workflows["in"]
	for {
		for _, rule := range currentWorkflow.rules {
			if rule.try(part) {
				if rule.target == "A" {
					return true
				} else if rule.target == "R" {
					return false
				}
				currentWorkflow = workflows[rule.target]
				break
			}
		}
	}
}

func partTwo(filename string) string {
	workflows, _ := parse(filename)

	currentWorkflow := workflows["in"]
	total := evaluateAcceptedPartsCount(workflows, currentWorkflow, map[string][]int{"x": {1, 4000}, "m": {1, 4000}, "a": {1, 4000}, "s": {1, 4000}})

	return strconv.Itoa(total)
}

func calculateTotal(limits map[string][]int) int {
	total := 1
	for _, v := range limits {
		total *= (v[1] - v[0] + 1)
	}
	// fmt.Println("Adding", total, "to total")
	return total
}

func evaluateAcceptedPartsCount(workflows map[string]Workflow, currentWorkflow Workflow, limits map[string][]int) int {

	total := 0
	for _, rule := range currentWorkflow.rules {
		// fmt.Println(currentWorkflow.name, ruleIdx, rule, limits)
		if rule.op == "<" {
			if limits[rule.category][0] < rule.value {
				limitsCopy := copyArrayMap(limits)
				limitsCopy[rule.category][1] = min(rule.value-1, limitsCopy[rule.category][1])
				limits[rule.category][0] = max(rule.value, limits[rule.category][0])

				// fmt.Println(limits, limitsCopy)

				if rule.target == "A" {
					total += calculateTotal(limitsCopy)
				} else if rule.target != "R" {
					total += evaluateAcceptedPartsCount(workflows, workflows[rule.target], limitsCopy)
				}
			}
		} else {
			if limits[rule.category][1] > rule.value {
				limitsCopy := copyArrayMap(limits)
				limitsCopy[rule.category][0] = max(rule.value+1, limitsCopy[rule.category][0])
				limits[rule.category][1] = min(rule.value, limits[rule.category][1])

				// fmt.Println(limits, limitsCopy)

				if rule.target == "A" {
					total += calculateTotal(limitsCopy)
				} else if rule.target != "R" {
					total += evaluateAcceptedPartsCount(workflows, workflows[rule.target], limitsCopy)
				}

				if rule.value == -1 {
					break
				}
			}
		}
	}
	return total
}
