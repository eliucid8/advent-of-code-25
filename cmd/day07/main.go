package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eliucid8/advent-of-code-2025/internal/util"
)

func part1(lines []string) string {
	var beams []bool
	for _, c := range lines[0] {
		beams = append(beams, c == 'S')
	}

	var num_splits int
	for _, row := range lines {
		for j, c := range row {
			if !beams[j] {
				continue
			}

			if c == '^' {
				num_splits++
				if j > 0 {
					beams[j-1] = true
				}
				if j < len(lines[0])-1 {
					beams[j+1] = true
				}

				beams[j] = false
			}
		}
	}
	return strconv.Itoa(num_splits)
}

func part2(lines []string) string {
	var beams = make(map[int]int)
	var num_splits uint64 = 1
	for i, c := range lines[0] {
		if c == 'S' {
			beams[i] = 1
		}
	}

	for _, row := range lines {
		for j, c := range row {
			if beams[j] < 1 {
				continue
			}

			if c == '^' {
				// fmt.Println(beams[j])
				num_splits += uint64(beams[j])
				if j > 0 {
					beams[j-1] += beams[j]
				}
				if j < len(lines[0])-1 {
					beams[j+1] += beams[j]
				}

				beams[j] = 0
			}

		}
		// fmt.Printf("=%d=\n", num_splits)
	}
	var total_paths int
	for _, count := range beams {
		total_paths += count
	}

	return strconv.FormatUint(num_splits, 10)
}

func main() {
	lines, err := util.ReadLines("data/day07/input.txt")
	if err != nil {
		log.Fatalf("read input: %v", err)
	}

	answer1 := part1(lines)
	fmt.Println("part 1:", answer1)
	answer2 := part2(lines)
	fmt.Println("part 2:", answer2)
}
