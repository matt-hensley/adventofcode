package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1: ", part1(input, "a"))
	fmt.Println("Part 2: ", part2(input))
}

func part1(input string, target string) int {
	return solve(input, target, nil)
}

func part2(input string) int {
	a := solve(input, "a", nil)
	return solve(input, "a", map[string]int{"b": a})
}

func solve(input string, target string, outputsOverride map[string]int) int {
	lines := strings.Split(input, "\n")
	ops := map[string][]string{}
	outputs := map[string]int{}

	if outputsOverride != nil {
		for k, v := range outputsOverride {
			outputs[k] = v
		}
	}

	for _, line := range lines {
		// example: y RSHIFT 2 -> g
		parts := strings.Split(line, " -> ")
		ops[parts[1]] = strings.Split(parts[0], " ")
	}

	var run func(string) int
	run = func(key string) int {
		// return if memoized i.e. already calculated
		if val, ok := outputs[key]; ok {
			return val
		}

		// return if constant
		if val, err := strconv.Atoi(key); err == nil {
			// memoize value before returning
			outputs[key] = val
			return val
		}

		op := ops[key]

		if len(op) == 1 {
			// constant or wire
			outputs[key] = run(op[0])
		} else if len(op) == 2 {
			// NOT
			outputs[key] = math.MaxUint16 ^ run(op[1])
		} else if len(op) == 3 {
			// AND, OR, LSHIFT, RSHIFT
			switch op[1] {
			case "AND":
				outputs[key] = run(op[0]) & run(op[2])
			case "OR":
				outputs[key] = run(op[0]) | run(op[2])
			case "LSHIFT":
				outputs[key] = run(op[0]) << run(op[2])
			case "RSHIFT":
				outputs[key] = run(op[0]) >> run(op[2])
			}
		} else {
			panic("unknown op")
		}

		return outputs[key]
	}

	return run(target)
}
