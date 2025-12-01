package main

import (
    "fmt"
    "log"

    "github.com/you/advent_of_code_go/internal/util"
)

func main() {
    lines, err := util.ReadLines("data/day02/input.txt")
    if err != nil {
        log.Fatalf("read input: %v", err)
    }
    fmt.Println("Line count:", len(lines))
    _ = lines
}
