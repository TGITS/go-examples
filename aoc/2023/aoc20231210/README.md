# AoC 2023 Day 10

**French version available:** [README_FR.md](README_FR.md)

## Overview

Go solution for [Advent of Code 2023 - Day 10](https://adventofcode.com/2023/day/10) : _Pipe Maze_

## Run

From this folder (`aoc/2023/aoc20231210`):

```powershell
go run .
```

CLI help:

```powershell
go run . -h
```

Available flags:

- `-demo-images`: generate only test preview PNG files
- `-pixel-size`: pixel size for PNG tile rendering (default: `2`)
- `-no-ascii`: disable ASCII visualizations (console and `.txt` files)
- `-png-only`: alias of `-no-ascii`

Common usage:

| Use case | Command |
|---|---|
| Standard run (ASCII + PNG) | `go run .` |
| PNG only (no ASCII output) | `go run . -png-only` |
| PNG only with bigger tiles | `go run . -png-only -pixel-size 4` |
| Generate only test preview PNGs | `go run . -demo-images` |
| Test previews with custom tile size | `go run . -demo-images -pixel-size 16` |
| Show CLI help | `go run . -h` |

## Test

From this folder (`aoc/2023/aoc20231210`):

```powershell
go test ./...
```

## Visualization

The solver now generates both ASCII and PNG visualizations.

### ASCII output

When running `go run .`, the program prints and saves:

- `visualization_loop.txt`
  - Loop-only view
  - Loop tiles use pipe characters
  - Non-loop tiles are `.`
- `visualization_distances.txt`
  - Loop distance view
  - Loop tiles use hexadecimal symbols (`0-9`, `A-F`) based on distance from start
  - Non-loop tiles are `.`

If `-no-ascii` or `-png-only` is set, this ASCII output is skipped.

### PNG output

When running `go run .`, the program also saves image visualizations:

- `visualization_loop_part1.png`
  - Part 1 rendering (loop vs non-loop)
- `visualization_enclosed_part2.png`
  - Part 2 rendering (loop + enclosed + outside)

### Color code

The PNG files use a simple color palette:

- **Blue** (`RGB 0,100,200`): tiles belonging to the main loop
- **Green** (`RGB 100,200,100`): tiles enclosed by the loop (Part 2 view)
- **Light gray** (`RGB 200,200,200`): tiles outside the loop

### Rendering model

- Each grid tile is rendered as a small filled square (`pixelSize` in code)
- Default CLI value is `-pixel-size 2`
- A larger pixel size can be used to produce more zoomed-in images

### Examples

Example below uses `./data/input_1_1.txt`.

- Loop view (`visualization_loop.txt` style):

```text
.....
.F-7.
.|.|.
.L-J.
.....
```

- Distance view (`visualization_distances.txt` style):

```text
.....
.012.
.1.3.
.234.
.....
```

- PNG legend reminder:
  - Blue: loop
  - Green: enclosed (Part 2 image)
  - Light gray: outside

- Example image files:
  - [Part 1 loop example](./test_loop_part1.png)
  - [Part 2 enclosed example](./test_enclosed_part2.png)

Part 1 loop preview:

<a href="./test_loop_part1.png">
  <img src="./test_loop_part1.png" alt="Part 1 loop preview" width="220" />
</a>

Part 2 enclosed preview:

<a href="./test_enclosed_part2.png">
  <img src="./test_enclosed_part2.png" alt="Part 2 enclosed preview" width="220" />
</a>

## Algorithm Explanation

### Problem Overview (Part 1)

You have pipes arranged in a two-dimensional grid of tiles

- the character `|` is a vertical pipe connecting north and south.
- the character `-` is a horizontal pipe connecting east and west.
- the character `L` is a 90-degree bend connecting north and east.
- the character `J` is a 90-degree bend connecting north and west.
- the character `7` is a 90-degree bend connecting south and west.
- the character `F` is a 90-degree bend connecting south and east.
- the character `.` is ground; there is no pipe in this tile.
- the character `S` is the starting position

There is a pipe on this tile `S`, but we do not directly know the shape this starting pipe has, we have to -deduce it-.
As a matter of fact, in the input data, we have to find the main loop, knowing that there are potentially pipes that aren't connected to this main loop.
To figure out which pipes form the main loop, we have to look to the ones connected to the starting pipe `S`, pipes those pipes connect to, pipes those pipes connect to, and so on. Every pipe in the main loop connects to its two neighbors, including S, which will have exactly two pipes connecting to it, and which is assumed to connect back to those two pipes.

When we have found the main loop, we need to find the tile that would take the longest number of steps along the loop to reach from the starting point regardless of which way around the loop we go.

The test input data are :

- `./data/input_1_1.txt`
  - The expected longest number of steps from the starting position is 4
- `./data/input_1_2.txt`
  - The expected longest number of steps from the starting position is 4
- `./data/input_1_3.txt`
  - The expected longest number of steps from the starting position is 8

The puzzle input data is `./data/input.txt` and we have to find the longest number of steps from the starting position for the first part of this puzzle.

### Solution Algorithm (Part 1)

To solve part 1 we implement the following steps:

1. **Parse grid**: read the input file line-by-line into a 2D slice of runes. Record the coordinates of the starting tile `S`.
2. **Determine start orientation**: the starting pipe may be any of the valid shapes. Inspect the four neighbouring cells (north, east, south, west) to see which are pipes (not `.`) and derive the two directions that the starting pipe must connect to. This determines an implicit shape for `S`.
3. **Walk the loop**: begin at the start coordinate and follow the pipe connections.
   - At each step, look at the current tile's character to know which exits are available (e.g. `-` only allows east/west, `L` allows north/east, etc.).
   - From the previous position, choose the next coordinate that is a valid exit and has not just been visited (do not go backwards).
   - Continue until returning to the start coordinate; during this traversal collect all visited coordinates in order.
4. **Compute distances**: the loop is circular. For each tile in the visited sequence compute its distance (step count) from the start assuming two directions around the cycle. The longer of the two distances is the number of steps required to reach that tile from the start.
5. **Answer**: the puzzle asks for the maximum of these distances over all loop tiles. Return that value.

This algorithm effectively performs a depth-first traversal constrained by the pipe connectivity rules, identifies the main loop by returning to `S`, and then measures the worst-case travel distance along the loop in either direction.

### Problem Overview (Part 2)

For part 2, we still use the same input data, with the same rules as in part 1 to define the pipes and the main loop.
However now we do not try to find the longest number of steps from the starting position.
For part 2 of the puzzle, we try to find the number of tiles enclosed by the loop.
There doesn't even need to be a full tile path to the outside for tiles to count as outside the loop : squeezing between pipes is also allowed, as in input data file `./data/input_2_2.txt`.
Any tile that isn't part of the main loop can count as being enclosed by the loop, as in input data file `./data/input_2_4.txt` : there are many bits of junk pipe lying around that aren't connected to the main loop at all.

The test input data are :

- `./data/input_2_1.txt`
  - The expected number of enclosed tiles is 4
- `./data/input_2_2.txt`
  - The expected number of enclosed tiles is 4
- `./data/input_2_3.txt`
  - The expected number of enclosed tiles is 8
- `./data/input_2_4.txt`
  - The expected number of enclosed tiles is 10

The puzzle input data is still `./data/input.txt` and we have to find the number of tiles enclosed by the loop for the second part of this puzzle.

### Solution Algorithm (Part 2)

To solve part 2, we use the **ray-casting algorithm** (also called the even-odd rule). The core idea is: a point is inside a polygon if a ray extending from that point crosses the polygon boundary an odd number of times.

#### Algorithm Steps:

1. **Find the main loop**: Using the same approach as Part 1, identify all coordinates that form the main loop pipe and store them in a set for fast lookup.

2. **For each tile not in the main loop**: Apply the ray-casting test
   - Cast a horizontal ray to the right from the tile (moving east from coordinates (x, y))
   - Count how many times this ray crosses the main loop boundary
   - A tile is *inside* the loop if the crossing count is **odd**
   - A tile is *outside* if the crossing count is **even**

3. **Handle pipe orientation edge cases**:
   - When the ray passes through a pipe character, we must consider the vertical extent of that pipe carefully to avoid double-counting at corners
   - Pipes `|` (vertical) always count as a boundary crossing when traversed horizontally
   - Pipes `-` (horizontal) do not count as a boundary crossing when traversed horizontally
   - Bend pipes (`L`, `J`, `7`, `F`) must be evaluated by checking the flow direction through the bend to determine if they represent a true boundary crossing or a tangent touch
   - A common technique: track entry and exit directions; if entry and exit are on opposite sides (north/south), count as a crossing; otherwise, it's a tangent

4. **Sum enclosed tiles**: Count all tiles where the ray-casting test indicates they are inside (odd parity).

#### Implementation Notes:

- The set of loop coordinates from Part 1 can be reused directly
- Care must be taken at pipe bends: for pipes like `L`, `J`, `7`, `F`, the direction of flow through them determines boundary behavior
- For efficient implementation: iterate through the grid once, and for each non-loop tile, count crossings by scanning rightward and checking loop membership
- Edge case: pipes at the edge of the grid need boundary checking to ensure the ray extends properly

