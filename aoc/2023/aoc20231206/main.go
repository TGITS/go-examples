package main

import (
	"log"

	"github.com/TGITS/go-examples/aoc/2023/aoc20231206/boatrace"
)

func main() {
	log.Printf("Product of record breakings for part 1: %d\n", boatrace.RecordBreakingsProducts("data/input.txt"))
	log.Printf("Number of ways to win for part 2: %d\n", boatrace.RecordBreaking("data/input.txt"))
}
