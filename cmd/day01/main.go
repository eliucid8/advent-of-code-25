package main

import (
    "fmt"
    "log"

    "github.com/you/advent_of_code_go/internal/util"
)

func main() {
    lines, err := util.ReadLines("data/day01/input.txt")
    if err != nil {
        log.Fatalf("read input: %v", err)
    }

    sum := 0
    for _, l := range lines {
        if l == "" {
            continue
        }
        var v int
        fmt.Sscan(l, &v)
        sum += v
    }

    fmt.Println("Sum:", sum)
}
