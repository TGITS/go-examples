package main

import (
	"log"

	"github.com/TGITS/go-examples/aoc/2023/aoc20231210/pipemaze"
)

func generateTestImages(pixelSize int) {
	log.Println("=== Test Input 1: Image Visualizations ===\n")

	// Part 1: Loop visualization
	log.Println("Part 1: Loop Visualization (Blue = Loop, Gray = Non-loop)")
	img1, err := pipemaze.GenerateLoopVisualizationImage("data/input_1_1.txt", pixelSize)
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		if err := pipemaze.SaveImageVisualization(img1, "test_loop_part1.png"); err != nil {
			log.Printf("Error: %v\n", err)
		} else {
			log.Println("✓ Test loop image saved to test_loop_part1.png\n")
		}
	}

	// Part 2: Enclosed visualization
	log.Println("Part 2: Enclosed Visualization (Blue = Loop, Green = Inside, Gray = Outside)")
	img2, err := pipemaze.GenerateEnclosedVisualizationImage("data/input_1_1.txt", pixelSize)
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		if err := pipemaze.SaveImageVisualization(img2, "test_enclosed_part2.png"); err != nil {
			log.Printf("Error: %v\n", err)
		} else {
			log.Println("✓ Test enclosed image saved to test_enclosed_part2.png")
		}
	}
}
