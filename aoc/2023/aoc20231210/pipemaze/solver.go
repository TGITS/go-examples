package pipemaze

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// dir represents a movement direction in the grid.
type dir struct{ dx, dy int }

// SolvePart1 reads the pipe maze from the given file and returns the
// maximum number of steps necessary to reach the farthest tile of the
// main loop from the starting position. Steps are counted along the loop
// and the shorter direction around the cycle is assumed (you may go either
// way).
func SolvePart1(filename string) int {
	grid, startX, startY := parseGrid(filename)
	visited := walkLoop(grid, startX, startY)
	return computeMaxDistance(visited)
}

// walkLoop traverses the pipe loop starting from (startX, startY) and returns
// the sequence of coordinates visited along the loop, in order.
func walkLoop(grid [][]rune, startX, startY int) [][2]int {
	neigh := startNeighbors(grid, startX, startY)
	if len(neigh) != 2 {
		log.Fatalf("start position has %d neighbours, expected 2", len(neigh))
	}

	// Begin walking by taking the first neighbor as the next tile
	prev := [2]int{startX - neigh[0].dx, startY - neigh[0].dy}
	current := [2]int{startX + neigh[0].dx, startY + neigh[0].dy}

	visited := make([][2]int, 0)
	visited = append(visited, [2]int{startX, startY})

	for {
		if current[0] == startX && current[1] == startY {
			break // closed the loop
		}
		visited = append(visited, current)

		next := findNext(grid, startX, startY, current, prev, visited)
		prev = current
		current = next
	}
	return visited
}

// findNext determines which tile to visit next given the current position and
// where we came from. It inspects the current tile's allowed exits and picks
// the one that leads to a valid neighbor (either another pipe or the start).
// It returns the next tile coordinates.
func findNext(grid [][]rune, startX, startY int, current, prev [2]int, visited [][2]int) [2]int {
	exits := allowedExits(grid[current[1]][current[0]])

	// Look for an exit that leads to a valid neighbor (not the previous tile)
	for _, d := range exits {
		nx := current[0] + d.dx
		ny := current[1] + d.dy

		if nx == prev[0] && ny == prev[1] {
			continue // Don't go backwards
		}

		// Only return to start if we've walked at least 3 tiles (start + 2 others)
		if nx == startX && ny == startY && len(visited) >= 3 {
			// Closed the loop
			return [2]int{nx, ny}
		}

		if neighborHasExit(grid, nx, ny, opposite(d)) {
			return [2]int{nx, ny}
		}
	}

	log.Fatalf("dead end at %v with char %c", current, grid[current[1]][current[0]])
	return [2]int{} // unreachable
}

// computeMaxDistance calculates the maximum of the minimal distances around
// a closed loop. For each position in the loop, it computes the distance from
// the start in both directions and takes the minimum. It returns the maximum
// of these minima.
func computeMaxDistance(visited [][2]int) int {
	length := len(visited)
	if length == 0 {
		return 0
	}

	maxMin := 0
	for i := range length {
		dist1 := i
		dist2 := length - i
		minDist := dist1
		if dist2 < minDist {
			minDist = dist2
		}
		if minDist > maxMin {
			maxMin = minDist
		}
	}
	return maxMin
}

// parseGrid reads the file and returns the grid as a slice of rune slices
// along with the coordinates of 'S' in the grid.
func parseGrid(filename string) ([][]rune, int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)

	startX, startY := -1, -1
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		for x, r := range row {
			if r == 'S' {
				startX, startY = x, y
			}
		}
		grid = append(grid, row)
		y++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return grid, startX, startY
}

// startNeighbors returns the directions from the start tile that lead to
// neighbouring pipe tiles that have an exit pointing back to the start.
func startNeighbors(grid [][]rune, sx, sy int) []dir {
	var out []dir
	for _, d := range []dir{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
		nx := sx + d.dx
		ny := sy + d.dy
		if ny >= 0 && ny < len(grid) && nx >= 0 && nx < len(grid[ny]) {
			if grid[ny][nx] != '.' {
				// Check that the neighbor has an exit pointing back to start
				if neighborHasExit(grid, nx, ny, opposite(d)) {
					out = append(out, d)
				}
			}
		}
	}
	return out
}

// directions allowed by a pipe tile character (excluding 'S').
func allowedExits(r rune) []dir {
	switch r {
	case '|':
		return []dir{{0, -1}, {0, 1}}
	case '-':
		return []dir{{-1, 0}, {1, 0}}
	case 'L':
		return []dir{{0, -1}, {1, 0}}
	case 'J':
		return []dir{{0, -1}, {-1, 0}}
	case '7':
		return []dir{{0, 1}, {-1, 0}}
	case 'F':
		return []dir{{0, 1}, {1, 0}}
	default:
		return nil
	}
}

// neighborHasExit checks that the tile at (x,y) has an exit in direction d.
func neighborHasExit(grid [][]rune, x, y int, d dir) bool {
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) {
		return false
	}
	exits := allowedExits(grid[y][x])
	for _, e := range exits {
		if e == d {
			return true
		}
	}
	return false
}

// opposite returns the reverse of a direction.
func opposite(d dir) dir {
	return dir{-d.dx, -d.dy}
}

// SolvePart2 counts the number of tiles enclosed by the main loop using ray-casting.
func SolvePart2(filename string) int {
	grid, startX, startY := parseGrid(filename)

	// Find the main loop
	visited := walkLoop(grid, startX, startY)

	// Convert visited to a set for O(1) lookup
	loopSet := make(map[[2]int]bool)
	for _, coord := range visited {
		loopSet[coord] = true
	}

	// Determine the actual character at the start position for ray-casting
	startChar := determineStartChar(grid, startX, startY)

	// Count enclosed tiles using ray-casting
	count := 0
	for y := range grid {
		for x := range grid[y] {
			if !loopSet[[2]int{x, y}] {
				// This tile is not part of the loop
				if isInside(grid, x, y, loopSet, startX, startY, startChar) {
					count++
				}
			}
		}
	}

	return count
}

// determineStartChar figures out what character the 'S' tile should be
// based on its neighboring connections.
func determineStartChar(grid [][]rune, sx, sy int) rune {
	neigh := startNeighbors(grid, sx, sy)

	// neigh contains the two directions from start that have valid connections
	// Map these directions to the actual pipe character
	has := make(map[dir]bool)
	for _, d := range neigh {
		has[d] = true
	}

	// Check all possible pipe characters to find which matches these directions
	for _, r := range []rune{'|', '-', 'L', 'J', '7', 'F'} {
		exits := allowedExits(r)
		if len(exits) != 2 {
			continue
		}
		// Check if this pipe's exits match our neighbors
		if has[exits[0]] && has[exits[1]] {
			return r
		}
	}

	log.Fatalf("Could not determine start character at (%d, %d)", sx, sy)
	return '?'
}

// isInside determines if a tile at (x, y) is inside the loop using ray-casting.
// It casts a horizontal ray to the right and counts boundary crossings.
func isInside(grid [][]rune, x, y int, loopSet map[[2]int]bool, startX, startY int, startChar rune) bool {
	crossings := 0
	var lastCorner rune

	// Cast ray to the right from (x, y) and count crossings
	for checkX := x + 1; checkX < len(grid[y]); checkX++ {
		if !loopSet[[2]int{checkX, y}] {
			continue // Not part of loop, skip
		}

		char := grid[y][checkX]

		// If this is the start position, use the determined character
		if checkX == startX && y == startY {
			char = startChar
		}

		switch char {
		case '|':
			// Vertical pipe always counts as a crossing
			crossings++

		case 'L', 'F':
			// Start of a corner sequence
			// L connects north and east, F connects south and east
			// When tracing left-to-right on a horizontal ray, we note the vertical extent
			lastCorner = char

		case 'J':
			// J connects north and west
			// If we came from F, we crossed from south to north = real crossing
			// If we came from L, we stayed at north = no crossing
			if lastCorner == 'F' {
				crossings++
			}

		case '7':
			// 7 connects south and west
			// If we came from L, we crossed from north to south = real crossing
			// If we came from F, we stayed at south = no crossing
			if lastCorner == 'L' {
				crossings++
			}

		case '-':
			// Horizontal pipe doesn't affect crossing count
		}
	}

	// Odd number of crossings = inside
	return crossings%2 == 1
}

// VisualizeLoopASCII creates a text visualization of the loop.
// Loop tiles are shown with their pipe character, non-loop tiles with '.'.
func VisualizeLoopASCII(filename string) string {
	grid, startX, startY := parseGrid(filename)
	visited := walkLoop(grid, startX, startY)

	// Convert visited to a set for O(1) lookup
	loopSet := make(map[[2]int]bool)
	for _, coord := range visited {
		loopSet[coord] = true
	}

	// Determine start character
	startChar := determineStartChar(grid, startX, startY)

	var result strings.Builder
	for y := range grid {
		for x := range grid[y] {
			if loopSet[[2]int{x, y}] {
				// Part of the loop - show the pipe character
				if x == startX && y == startY {
					result.WriteRune(startChar)
				} else {
					result.WriteRune(grid[y][x])
				}
			} else {
				// Not part of loop - show as space or dot
				result.WriteRune('.')
			}
		}
		result.WriteRune('\n')
	}

	return result.String()
}

// VisualizeLoopWithDistances creates a visualization showing distances from start.
// Loop tiles show their distance using hexadecimal notation (0-9, A-F, repeating).
// This provides a clean repeating pattern that stays within ASCII range.
// Non-loop tiles are shown as '.'.
func VisualizeLoopWithDistances(filename string) string {
	grid, startX, startY := parseGrid(filename)
	visited := walkLoop(grid, startX, startY)

	// Convert visited to a set and create distance map
	loopSet := make(map[[2]int]bool)
	distances := make(map[[2]int]int)
	for i, coord := range visited {
		loopSet[coord] = true
		// Distance is minimum from start in either direction around the loop
		dist1 := i
		dist2 := len(visited) - i
		minDist := dist1
		if dist2 < minDist {
			minDist = dist2
		}
		distances[coord] = minDist
	}

	var result strings.Builder
	for y := range grid {
		for x := range grid[y] {
			if loopSet[[2]int{x, y}] {
				// Part of the loop - show distance in hexadecimal (mod 16)
				dist := distances[[2]int{x, y}]
				hexDigit := dist % 16
				var char rune
				if hexDigit < 10 {
					char = rune('0' + hexDigit)
				} else {
					char = rune('A' + hexDigit - 10)
				}
				result.WriteRune(char)
			} else {
				// Not part of loop
				result.WriteRune('.')
			}
		}
		result.WriteRune('\n')
	}

	return result.String()
}

// SaveVisualization saves the ASCII visualization to a file.
func SaveVisualization(filename string, outputFile string, includeDistances bool) error {
	var visualization string
	if includeDistances {
		visualization = VisualizeLoopWithDistances(filename)
	} else {
		visualization = VisualizeLoopASCII(filename)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(visualization)
	return err
}

// PrintVisualization prints the ASCII visualization to stdout.
func PrintVisualization(filename string, includeDistances bool) {
	var visualization string
	if includeDistances {
		visualization = VisualizeLoopWithDistances(filename)
	} else {
		visualization = VisualizeLoopASCII(filename)
	}

	fmt.Print(visualization)
}
