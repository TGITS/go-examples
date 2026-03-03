package main

import (
	"flag"
	"log"

	"github.com/TGITS/go-examples/aoc/2023/aoc20231210/pipemaze"
)

func main() {
	demoImages := flag.Bool("demo-images", false, "Generate only test preview PNG files")
	pixelSize := flag.Int("pixel-size", 2, "Pixel size for PNG tile rendering (must be >= 1)")
	noASCII := flag.Bool("no-ascii", false, "Disable ASCII visualizations (console and .txt files), generate PNG only")
	pngOnly := flag.Bool("png-only", false, "Alias of -no-ascii: generate PNG visualizations without ASCII output")
	flag.Parse()

	if *pixelSize < 1 {
		log.Fatalf("invalid -pixel-size=%d: must be >= 1", *pixelSize)
	}

	if *noASCII && !*pngOnly {
		*pngOnly = true
	}
	if *pngOnly && !*noASCII {
		*noASCII = true
	}

	if *demoImages {
		generateTestImages(*pixelSize)
		return
	}

	log.Printf("Longest loop distance (Part 1): %d\n", pipemaze.SolvePart1("data/input.txt")) // 6823 for my input data
	log.Printf("Enclosed tiles (Part 2): %d\n", pipemaze.SolvePart2("data/input.txt"))        // 415 for my input data

	if !*noASCII {
		// Visualizations for Part 1 (ASCII)
		log.Println("\n=== Part 1: Loop Visualization (ASCII) ===")
		pipemaze.PrintVisualization("data/input.txt", false)

		log.Println("\n=== Part 1: Loop Visualization (with Distances) ===")
		pipemaze.PrintVisualization("data/input.txt", true)

		// Save ASCII visualizations to files
		if err := pipemaze.SaveVisualization("data/input.txt", "visualization_loop.txt", false); err != nil {
			log.Printf("Error saving loop visualization: %v\n", err)
		} else {
			log.Println("Loop visualization saved to visualization_loop.txt")
		}

		if err := pipemaze.SaveVisualization("data/input.txt", "visualization_distances.txt", true); err != nil {
			log.Printf("Error saving distance visualization: %v\n", err)
		} else {
			log.Println("Distance visualization saved to visualization_distances.txt")
		}
	} else {
		log.Println("ASCII visualizations disabled by -no-ascii/-png-only")
	}

	// Image visualizations for Part 1
	log.Println("\n=== Part 1: Image Visualization (Loop) ===")
	img1, err := pipemaze.GenerateLoopVisualizationImage("data/input.txt", *pixelSize)
	if err != nil {
		log.Printf("Error generating loop image: %v\n", err)
	} else {
		if err := pipemaze.SaveImageVisualization(img1, "visualization_loop_part1.png"); err != nil {
			log.Printf("Error saving loop image: %v\n", err)
		} else {
			log.Printf("Loop image saved to visualization_loop_part1.png (pixel-size=%d)\n", *pixelSize)
		}
	}

	// Image visualizations for Part 2
	log.Println("=== Part 2: Image Visualization (Enclosed) ===")
	img2, err := pipemaze.GenerateEnclosedVisualizationImage("data/input.txt", *pixelSize)
	if err != nil {
		log.Printf("Error generating enclosed image: %v\n", err)
	} else {
		if err := pipemaze.SaveImageVisualization(img2, "visualization_enclosed_part2.png"); err != nil {
			log.Printf("Error saving enclosed image: %v\n", err)
		} else {
			log.Printf("Enclosed image saved to visualization_enclosed_part2.png (pixel-size=%d)\n", *pixelSize)
		}
	}

}
