package main

import (
	"log"

	"github.com/TGITS/go-examples/aoc/2023/aoc20231207/camelcards"
)

func main() {
	log.Printf("Total winnings for part 1: %d\n", camelcards.SolvePart1("data/input.txt"))
	log.Printf("Total winnings for part 2: %d\n", camelcards.SolvePart2("data/input.txt"))
}
