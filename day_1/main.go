package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/aaronjprice/aoc_2024/lib"
)

func main() {
	solution, err := SolvePart1File("day_1/input.txt")
	lib.IfErrLogFatal(err)
	fmt.Println("part 1 solution:", solution)

	solution, err = SolvePart2File("day_1/input.txt")
	lib.IfErrLogFatal(err)
	fmt.Println("part 2 solution:", solution)
}

// Read file data and solve the puzzle
func SolvePart1File(filename string) (int, error) {
	ls, err := readLists(filename)
	if err != nil {
		return 0, fmt.Errorf("solving for file '%v': %w", filename, err)
	}
	return solvePart1(ls), nil
}

type lists struct {
	left  []int
	right []int
}

// Read file with filename and parse contents into the lists type
func readLists(filename string) (lists, error) {
	lines, err := lib.ReadLines(filename)
	if err != nil {
		return lists{}, err
	}

	length := len(lines)
	output := lists{
		left:  make([]int, length),
		right: make([]int, length),
	}

	for i, line := range lines {
		split := strings.Split(line, "   ")
		if len(split) != 2 {
			return output, fmt.Errorf("reading lists: line [%v] '%v': splitting on 3 spaces '   ' did not generate two strings", i, line)
		}
		left, err := strconv.Atoi(split[0])
		if err != nil {
			return output, fmt.Errorf("reading lists: line [%v] '%v': converting left: %w", i, line, err)
		}
		right, err := strconv.Atoi(split[1])
		if err != nil {
			return output, fmt.Errorf("reading lists: line [%v] '%v': converting right: %w", i, line, err)
		}
		output.left[i] = left
		output.right[i] = right
	}

	return output, nil
}

// Solve part 1 of the puzzle
func solvePart1(ls lists) int {
	slices.Sort(ls.left)
	slices.Sort(ls.right)

	total := 0
	for i := 0; i < len(ls.left); i++ {
		diff := abs(ls.left[i] - ls.right[i])
		total += diff
	}

	return total
}

// Get the absolute value of an integer
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// Read file data and solve the puzzle
func SolvePart2File(filename string) (int, error) {
	ls, err := readLists(filename)
	if err != nil {
		return 0, fmt.Errorf("solving for file '%v': %w", filename, err)
	}
	return solvePart2(ls), nil
}

// Solve part 2 of the puzzle
func solvePart2(ls lists) int {
	score := 0

	for _, l := range ls.left {
		for _, r := range ls.right {
			if l == r {
				score += l
			}
		}
	}

	return score
}
