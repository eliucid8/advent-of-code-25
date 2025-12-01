package main

import (
    "fmt"
    "log"
	"strconv"

    "github.com/eliucid8/advent-of-code-25/internal/util"
)

func main() {
    lines, err := util.ReadLines("data/day01/input.txt")
    if err != nil {
        log.Fatalf("read input: %v", err)
    }

    displacement := 0
	sum := 0
    for _, l := range lines {
        if l == "" {
            continue
        }
		distance, _ := strconv.Atoi(l[1:])
		if l[0] == "R" {
			displacement += distance
		} else {
			displacement -= distance
		}
	    displacement %= 100
		if displacement == 0 {
			sum += 1
		}
    }

    fmt.Println("sum:", sum)
}
