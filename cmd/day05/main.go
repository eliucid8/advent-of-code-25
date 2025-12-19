package main

import (
	"cmp"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/eliucid8/advent-of-code-2025/internal/util"
)

func part1(lines []string) string {
	var ranges [][]uint64
	var i = 0
	for lines[i] != "" {
		var freshRange = strings.Split(lines[i], "-")
		var low, _ = strconv.ParseUint(freshRange[0], 10, 64)
		var high, _ = strconv.ParseUint(freshRange[1], 10, 64)

		ranges = append(ranges, []uint64{low, high})
		i++
	}
	i++
	var ingredients []uint64
	for i < len(lines) {
		var parsed, _ = strconv.ParseUint(lines[i], 10, 64)
		ingredients = append(ingredients, parsed)
		i++
	}

	// sort slices, place least valued one in the back
	slices.Sort(ingredients)
	slices.Reverse(ingredients)

	// sort fresh ranges by bottom of range
	slices.SortFunc(ranges, func(a, b []uint64) int {
		return cmp.Compare(a[0], b[0])
	})

	fresh := 0
	for _, freshRange := range ranges {
		// start at back of ingredients
		i := len(ingredients) - 1
		// while ingredient is less than the top of the range
		for ingredients[i] <= freshRange[1] {
			if ingredients[i] >= freshRange[0] {
				fresh++
			}
			// remove this element from ingredients
			ingredients = ingredients[:len(ingredients)-1]
			i--
		}
	}

	return strconv.Itoa(fresh)
}

func part2(lines []string) string {
	var freshRanges [][]uint64
	var i = 0
	for lines[i] != "" {
		var freshRange = strings.Split(lines[i], "-")
		var low, _ = strconv.ParseUint(freshRange[0], 10, 64)
		var high, _ = strconv.ParseUint(freshRange[1], 10, 64)

		freshRanges = append(freshRanges, []uint64{low, high})
		i++
	}

	// sort fresh ranges by bottom of range
	slices.SortFunc(freshRanges, func(a, b []uint64) int {
		return cmp.Compare(a[0], b[0])
	})

	// quick little interval merge >:)
	// this is the problem that got me my job lmao
	var newRanges [][]uint64
	var curRange = freshRanges[0]
	for _, freshRange := range freshRanges {
		if freshRange[0] > curRange[1] {
			newRanges = append(newRanges, curRange)
			curRange = freshRange
		} else {
			curRange[1] = max(curRange[1], freshRange[1])
		}
	}
	newRanges = append(newRanges, curRange)

	// guaranteed to be disjoint now i think
	var total uint64 = 0
	for _, freshRange := range newRanges {
		total += freshRange[1] - freshRange[0] + 1
	}

	return strconv.FormatUint(total, 10)
}

func main() {
	lines, err := util.ReadLines("data/day05/input.txt")
	if err != nil {
		log.Fatalf("read input: %v", err)
	}

	answer1 := part1(lines)
	fmt.Println("part 1:", answer1)
	answer2 := part2(lines)
	fmt.Println("part 2:", answer2)
}
