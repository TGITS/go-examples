package main

import (
	"log"

	"github.com/TGITS/go-examples/aoc/2023/aoc20231205/almanac"
)

// main runs both puzzle parts against the full input file.
func main() {
	log.Printf("minimum location for part 1: %d\n", almanac.GetMinimumLocationForPart1("data/input.txt"))
	log.Printf("minimum location for part 2: %d\n", almanac.GetMinimumLocationForPart2("data/input.txt"))
}
