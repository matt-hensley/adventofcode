package main

import (
	"crypto/md5"
	_ "embed"
	"fmt"
	"io"
  "strconv"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}

func part1(input string) int {
	match := false
  num := 0

	for match != true {
		h := md5.New()
		io.WriteString(h, input)
		io.WriteString(h, strconv.Itoa(num))
    sum := fmt.Sprintf("%x", h.Sum(nil))
    
    if sum[:5] == "00000" {
      return num
    }

    num++
	}
  
	return -1
}

func part2(input string) int {
	match := false
  num := 0

	for match != true {
		h := md5.New()
		io.WriteString(h, input)
		io.WriteString(h, strconv.Itoa(num))
    sum := fmt.Sprintf("%x", h.Sum(nil))
    
    if sum[:6] == "000000" {
      return num
    }

    num++
	}
  
	return -1
}
