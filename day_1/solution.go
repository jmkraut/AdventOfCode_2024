package main

import (
	"slices"
	"sort"
	"strconv"
	"strings"
)

func SolvePart1(lines []string) int {
	leftSlice, rightSlice := stringSliceToDualIntSlice(lines)

	slices.Sort(sort.IntSlice(leftSlice))
	slices.Sort(sort.IntSlice(rightSlice))

	if len(leftSlice) != len(rightSlice) {
		panic("arrays are not of the same length")
	}

	var totalDistance int
	for i := range leftSlice {
		var distance int

		if rightSlice[i] > leftSlice[i] {
			distance = rightSlice[i] - leftSlice[i]
		} else {
			distance = leftSlice[i] - rightSlice[i]
		}

		totalDistance += distance
	}

	return totalDistance
}

func SolvePart2(lines []string) int {
	return 0
}

func stringSliceToDualIntSlice(lines []string) ([]int, []int) {
	var leftSlice, rightSlice []int

	for i := range lines {
		splitString := strings.Split(lines[i], "   ")
		leftInt, _ := strconv.Atoi(splitString[0])
		rightInt, _ := strconv.Atoi(splitString[1])
		leftSlice = append(leftSlice, leftInt)
		rightSlice = append(rightSlice, rightInt)
	}

	return leftSlice, rightSlice
}
