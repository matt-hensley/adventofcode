package main

import (
  _ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}

func part1(input string) int {
  reject := map[string]bool {
    "ab": true,
    "cd": true,
    "pq": true,
    "xy": true,
  }
  
  lines := strings.Split(input, "\n")
  nice := 0

  for _, line := range lines {
    disallowed := false
    vowels := 0
    doubles := 0

    // obviously a tasks for regexes but this is more interesting
    for i := 0; i < len(line); i++ {
      if i < len(line) - 1 {
        pair := line[i:i+2]
  
        // map lookups are fast?
        if reject[pair] == true {
          disallowed = true
          break
        }
  
        if pair[0] == pair[1] {
          doubles++
        }
      }

      c := line[i:i+1]

      if strings.Contains("aeiou", c) {
        vowels++
      }
    }

    if !disallowed && vowels >= 3 && doubles >= 1 {
      nice++
    }
  }
  
	return nice
}

func part2(input string) int {
  lines := strings.Split(input, "\n")
  nice := 0

  for _, line := range lines {
    paired := false
    allowed_triple := false
    
    for i := 0; i < len(line) - 1; i++ {
      paired = strings.Contains(line[i+2:], line[i:i+2])

      if paired {
        break
      }
    }

    if !paired {
      continue
    }
    
    for i := 0; i < len(line) - 2; i++ {
      triple := line[i:i+3]

      if triple[0] == triple[2] {
        allowed_triple = true
        break
      }
    }

    if paired == true && allowed_triple == true {
      nice++
    }
    
  }
  
	return nice
}
