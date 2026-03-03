package pipemaze

import "testing"

func TestSolvePart1_WithExamples(t *testing.T) {
	cases := []struct {
		filename string
		expected int
	}{
		{"../data/input_1_1.txt", 4},
		{"../data/input_1_2.txt", 4},
		{"../data/input_1_3.txt", 8},
	}

	for _, c := range cases {
		actual := SolvePart1(c.filename)
		if actual != c.expected {
			t.Fatalf("%s: expected %d, got %d", c.filename, c.expected, actual)
		}
	}
}

// ==================== Intermediate Function Tests ====================

func TestAllowedExits(t *testing.T) {
	cases := []struct {
		char     rune
		expected []dir
	}{
		{'|', []dir{{0, -1}, {0, 1}}},
		{'-', []dir{{-1, 0}, {1, 0}}},
		{'L', []dir{{0, -1}, {1, 0}}},
		{'J', []dir{{0, -1}, {-1, 0}}},
		{'7', []dir{{0, 1}, {-1, 0}}},
		{'F', []dir{{0, 1}, {1, 0}}},
		{'.', nil},
		{'S', nil},
	}

	for _, c := range cases {
		actual := allowedExits(c.char)
		if !dirsEqual(actual, c.expected) {
			t.Fatalf("allowedExits('%c'): expected %v, got %v", c.char, c.expected, actual)
		}
	}
}

func TestOpposite(t *testing.T) {
	cases := []struct {
		input    dir
		expected dir
	}{
		{dir{0, -1}, dir{0, 1}},
		{dir{0, 1}, dir{0, -1}},
		{dir{1, 0}, dir{-1, 0}},
		{dir{-1, 0}, dir{1, 0}},
	}

	for _, c := range cases {
		actual := opposite(c.input)
		if actual != c.expected {
			t.Fatalf("opposite(%v): expected %v, got %v", c.input, c.expected, actual)
		}
	}
}

func TestComputeMaxDistance(t *testing.T) {
	cases := []struct {
		visited  [][2]int
		expected int
	}{
		{
			// Simple 4-tile loop: start -> a -> b -> c -> start
			// Distances: start=0, a=1, b=2, c=1, start=0
			// Max min distance: max(0, min(1,3), min(2,2), min(1,3)) = max(0,1,2,1) = 2
			[][2]int{{0, 0}, {1, 0}, {1, 1}, {0, 1}},
			2,
		},
		{
			// Single tile (shouldn't happen, but test edge case)
			[][2]int{{0, 0}},
			0,
		},
		{
			// Empty (edge case)
			[][2]int{},
			0,
		},
		{
			// Simple 8-tile loop
			[][2]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}, {4, 3}},
			4,
		},
	}

	for i, c := range cases {
		actual := computeMaxDistance(c.visited)
		if actual != c.expected {
			t.Fatalf("computeMaxDistance (case %d): expected %d, got %d", i, c.expected, actual)
		}
	}
}

func TestParseGrid(t *testing.T) {
	grid, startX, startY := parseGrid("../data/input_1_1.txt")

	// Test grid dimensions
	if len(grid) != 5 {
		t.Fatalf("grid rows: expected 5, got %d", len(grid))
	}
	if len(grid[0]) != 5 {
		t.Fatalf("grid cols: expected 5, got %d", len(grid[0]))
	}

	// Test start position
	if startX != 1 || startY != 1 {
		t.Fatalf("start position: expected (1, 1), got (%d, %d)", startX, startY)
	}

	// Test specific characters
	if grid[1][1] != 'S' {
		t.Fatalf("start char: expected 'S', got '%c'", grid[1][1])
	}
	if grid[1][2] != '-' {
		t.Fatalf("char at (2,1): expected '-', got '%c'", grid[1][2])
	}
}

func TestStartNeighbors(t *testing.T) {
	grid, startX, startY := parseGrid("../data/input_1_1.txt")
	neighbors := startNeighbors(grid, startX, startY)

	if len(neighbors) != 2 {
		t.Fatalf("expected 2 neighbors, got %d", len(neighbors))
	}

	// For input_1_1.txt, S at (1,1) connects to:
	// - right: (2,1) which is '-' (direction (1, 0))
	// - down:  (1,2) which is '|' (direction (0, 1))
	hasEast := false
	hasSouth := false
	for _, n := range neighbors {
		if n == (dir{1, 0}) {
			hasEast = true
		}
		if n == (dir{0, 1}) {
			hasSouth = true
		}
	}
	if !hasEast || !hasSouth {
		t.Fatalf("expected east and south neighbors, got %v", neighbors)
	}
}

func TestWalkLoop_Simple(t *testing.T) {
	grid, startX, startY := parseGrid("../data/input_1_1.txt")
	visited := walkLoop(grid, startX, startY)

	// input_1_1.txt forms an 8-tile loop:
	// S(1,1) -> -(2,1) -> 7(3,1) -> |(3,2) -> J(3,3) -> -(2,3) -> L(1,3) -> |(1,2) -> S
	if len(visited) != 8 {
		t.Fatalf("expected 8 visited tiles, got %d\nvisited: %v", len(visited), visited)
	}

	// Start should be first
	if visited[0][0] != startX || visited[0][1] != startY {
		t.Fatalf("first visited should be start (%d, %d), got (%d, %d)", startX, startY, visited[0][0], visited[0][1])
	}
}

// dirsEqual checks if two slices of directions are equal (order-independent).
func dirsEqual(a, b []dir) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// ==================== Part 2 Tests ====================

func TestSolvePart2_WithExamples(t *testing.T) {
	cases := []struct {
		filename string
		expected int
	}{
		{"../data/input_2_1.txt", 4},
		{"../data/input_2_2.txt", 4},
		{"../data/input_2_3.txt", 8},
		{"../data/input_2_4.txt", 10},
	}

	for _, c := range cases {
		actual := SolvePart2(c.filename)
		if actual != c.expected {
			t.Fatalf("%s: expected %d, got %d", c.filename, c.expected, actual)
		}
	}
}

func TestDetermineStartChar(t *testing.T) {
	cases := []struct {
		filename string
		expected rune
	}{
		{"../data/input_1_1.txt", 'F'}, // At (1,1): connects east and south
		{"../data/input_1_2.txt", 'F'}, // At (1,1): connects east and south
		{"../data/input_1_3.txt", 'F'}, // At (1,1): connects east and south
	}

	for _, c := range cases {
		grid, startX, startY := parseGrid(c.filename)
		actual := determineStartChar(grid, startX, startY)
		if actual != c.expected {
			t.Fatalf("%s: expected '%c', got '%c'", c.filename, c.expected, actual)
		}
	}
}
