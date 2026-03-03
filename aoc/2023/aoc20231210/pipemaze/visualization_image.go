package pipemaze

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// GenerateLoopVisualizationImage creates a PNG image for Part 1.
// Blue pixels: loop tiles, Gray pixels: non-loop tiles
func GenerateLoopVisualizationImage(filename string, pixelSize int) (image.Image, error) {
	grid, startX, startY := parseGrid(filename)
	visited := walkLoop(grid, startX, startY)

	// Convert visited to a set
	loopSet := make(map[[2]int]bool)
	for _, coord := range visited {
		loopSet[coord] = true
	}

	height := len(grid)
	width := 0
	if height > 0 {
		width = len(grid[0])
	}

	// Create image
	imgWidth := width * pixelSize
	imgHeight := height * pixelSize
	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	// Colors
	colorLoop := color.RGBA{0, 100, 200, 255}      // Blue
	colorOutside := color.RGBA{200, 200, 200, 255} // Light gray

	// Draw pixels
	for y := range height {
		for x := range width {
			var c color.Color
			if loopSet[[2]int{x, y}] {
				c = colorLoop
			} else {
				c = colorOutside
			}

			// Draw pixelSize x pixelSize block
			for dy := range pixelSize {
				for dx := range pixelSize {
					img.Set(x*pixelSize+dx, y*pixelSize+dy, c)
				}
			}
		}
	}

	return img, nil
}

// GenerateEnclosedVisualizationImage creates a PNG image for Part 2.
// Blue pixels: loop tiles, Green pixels: enclosed tiles, Gray pixels: outside
func GenerateEnclosedVisualizationImage(filename string, pixelSize int) (image.Image, error) {
	grid, startX, startY := parseGrid(filename)
	visited := walkLoop(grid, startX, startY)

	// Convert visited to a set for O(1) lookup
	loopSet := make(map[[2]int]bool)
	for _, coord := range visited {
		loopSet[coord] = true
	}

	// Determine start character
	startChar := determineStartChar(grid, startX, startY)

	height := len(grid)
	width := 0
	if height > 0 {
		width = len(grid[0])
	}

	// Create image
	imgWidth := width * pixelSize
	imgHeight := height * pixelSize
	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	// Colors
	colorLoop := color.RGBA{0, 100, 200, 255}      // Blue
	colorInside := color.RGBA{100, 200, 100, 255}  // Green
	colorOutside := color.RGBA{200, 200, 200, 255} // Light gray

	// Draw pixels
	for y := range height {
		for x := range width {
			var c color.Color

			if loopSet[[2]int{x, y}] {
				c = colorLoop
			} else if isInside(grid, x, y, loopSet, startX, startY, startChar) {
				c = colorInside
			} else {
				c = colorOutside
			}

			// Draw pixelSize x pixelSize block
			for dy := range pixelSize {
				for dx := range pixelSize {
					img.Set(x*pixelSize+dx, y*pixelSize+dy, c)
				}
			}
		}
	}

	return img, nil
}

// SaveImageVisualization saves an image to a PNG file.
func SaveImageVisualization(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}
