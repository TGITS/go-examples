package main

import (
	"fmt"
	"log"

	"github.com/TGITS/go-examples/aoc/2023/aoc20231210/pipemaze"
)

// runTestVisualizationDemo generates and saves visualizations for test input
func runTestVisualizationDemo() {
	fmt.Println("\n=== Test Input 1: Loop Visualization ===")
	pipemaze.PrintVisualization("data/input_1_1.txt", false)
	fmt.Println("\n=== Test Input 1: Distance Visualization ===")
	pipemaze.PrintVisualization("data/input_1_1.txt", true)

	// Save test visualizations
	if err := pipemaze.SaveVisualization("data/input_1_1.txt", "test_visualization_loop.txt", false); err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Println("\nTest loop visualization saved to test_visualization_loop.txt")
	}

	if err := pipemaze.SaveVisualization("data/input_1_1.txt", "test_visualization_distances.txt", true); err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Test distance visualization saved to test_visualization_distances.txt")
	}
}
