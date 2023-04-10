package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Op int

const (
	Toggle Op = 1
	On        = 2
	Off       = 3
)

type Point struct{ x, y int }

type Callback = func(x int) int

func main() {
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}

func part1(input string) int {
	// experiment to see if Go can work more like JS or Python
	commands := map[Op]Callback{
		Toggle: func(x int) int {
			if x == 1 {
				return 0
			}

			return 1
		},
		On:  func(_ int) int { return 1 },
		Off: func(_ int) int { return 0 },
	}

	lines := strings.Split(input, "\n")
	var lights [1000][1000]int

	for _, line := range lines {
		parts := strings.Split(line, " ")
		var op Op
		var start, stop Point

		switch parts[0] {
		case "toggle":
			op = Toggle
			start = string_to_point(parts[1])
			stop = string_to_point(parts[3])
		case "turn":
			switch parts[1] {
			case "on":
				op = On
			case "off":
				op = Off
			}

			start = string_to_point(parts[2])
			stop = string_to_point(parts[4])
		default:
			panic("Parsing failure")
		}

		for x := start.x; x <= stop.x; x++ {
			for y := start.y; y <= stop.y; y++ {
				lights[x][y] = commands[op](lights[x][y])
			}
		}
	}

	count := 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			count += lights[i][j]
		}
	}

	return count
}

func part2(input string) int {
	commands := map[Op]Callback{
		Toggle: func(x int) int { return x + 2 },
		On:     func(x int) int { return x + 1 },
		Off: func(x int) int {
			if x == 0 {
				return 0
			}
			return x - 1
		},
	}

	lines := strings.Split(input, "\n")
	var lights [1000][1000]int

	for _, line := range lines {
		parts := strings.Split(line, " ")
		var op Op
		var start, stop Point

		switch parts[0] {
		case "toggle":
			op = Toggle
			start = string_to_point(parts[1])
			stop = string_to_point(parts[3])
		case "turn":
			switch parts[1] {
			case "on":
				op = On
			case "off":
				op = Off
			}

			start = string_to_point(parts[2])
			stop = string_to_point(parts[4])
		default:
			panic("Parsing failure")
		}

		for x := start.x; x <= stop.x; x++ {
			for y := start.y; y <= stop.y; y++ {
				lights[x][y] = commands[op](lights[x][y])
			}
		}
	}

	count := 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			count += lights[i][j]
		}
	}

	return count
}

func string_to_point(pair string) Point {
	var x int
	var y int
	fmt.Sscanf(pair, "%d,%d", &x, &y)
	return Point{x, y}
}
