package main

import (
	"fmt"
	"log"

	"github.com/jmkraut/aoc2024/input"
)

func main() {
	lines, err := input.ReadLines("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	part1Result := SolvePart1(lines)
	fmt.Printf("Part 1: %v\n", part1Result)

	part2Result := SolvePart2(lines)
	fmt.Printf("Part 2: %v\n", part2Result)
}
