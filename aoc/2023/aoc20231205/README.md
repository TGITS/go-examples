# AoC 2023 Day 5

**French version available:** [README_FR.md](README_FR.md)

## Overview

Go solution for [Advent of Code 2023 - Day 5](https://adventofcode.com/2023/day/5)

## Run

From this folder (`aoc/2023/aoc20231205`):

```powershell
go run .
```

## Test

From this folder (`aoc/2023/aoc20231205`):

```powershell
go test ./...
```

## Benchmarks

This module includes dedicated benchmarks to compare:
- `BruteForce` (seed-by-seed original part 2 approach)
- `Optimized` (interval pipeline approach)

From this folder (`aoc/2023/aoc20231205`), run:

```powershell
# Optimized on full input (single run)
go test ./almanac -bench "BenchmarkPart2Optimized_Input$" -benchmem -benchtime=1x -run "^$"

# Brute force on full input (single run, may be very long)
go test ./almanac -bench "BenchmarkPart2BruteForce_Input$" -benchmem -benchtime=1x -run "^$"

# Quick comparison on test input
go test ./almanac -bench "BenchmarkPart2(Optimized|BruteForce)_InputTest$" -benchmem -benchtime=1x -run "^$"
```

Notes:
- Benchmarks are "clean": puzzle parsing logs are silenced during benchmark setup.
- Running full-input optimized and brute-force in separate commands avoids incomplete output when brute force is interrupted.

Or run the helper script:

```powershell
powershell -ExecutionPolicy Bypass -File ".\run_benchmarks.ps1"

# Skip brute force on full input
powershell -ExecutionPolicy Bypass -File ".\run_benchmarks.ps1" -SkipBruteForce
```

Equivalent Bash script:

```bash
bash ./run_benchmarks.sh

# Skip brute force on full input
bash ./run_benchmarks.sh --skip-bruteforce
```

## Algorithm Explanation

### Problem Overview

This puzzle requires mapping seed IDs through a series of transformations (seed → soil → fertilizer → water → light → temperature → humidity → location) using interval-based mappings defined in an "Almanac". The goal is to find the minimum location number across all seeds.

### Part 1: Individual Seed Mapping (Brute Force)

**Approach:**
The naive solution treats each seed separately and chains all transformation functions:

1. Parse seeds as individual values
2. For each seed, apply the transformation pipeline:
   - Look up which mapping rule (if any) applies to the current value
   - If found, translate the value using: `destination + (value - source)`
   - If no rule applies, the value maps to itself (identity mapping)
   - Move to the next transformation stage
3. Collect all final location values and return the minimum

**Time Complexity:** O(n × m) where n is the number of seeds and m is the total number of mapping rules across all stages.

**Limitation:** Part 2 specifies seed ranges (pairs of start + length), which can represent billions of individual seeds, making this approach impractical.

### Part 2: Interval Pipeline (Optimized)

**Key Insight:** Rather than iterating each individual seed, we can propagate ranges of seeds through the mappings as intervals. Intervals with identical transformation behavior can be grouped and processed as a single unit.

**Data Structure:**
- **Interval:** A half-open range `[start, end)` representing contiguous seed/value ranges
- **Half-open ranges** simplify arithmetic and merging: a range always excludes its end point, avoiding off-by-one errors

**Algorithm Steps:**

1. **Parse Seed Ranges:**
   - Interpret seed data as pairs: `(start, length)` → interval `[start, start + length)`
   - Example: seeds `79 14 55 13` become intervals `[79, 93)` and `[55, 68)`

2. **Normalize Intervals:**
   - Sort intervals by start position
   - Merge overlapping or adjacent intervals to reduce fragmentation
   - Example: `[79, 93)` and `[88, 100)` merge into `[79, 100)`

3. **Apply Mapping to Intervals (`applyMapToIntervals`):**
   For each input interval, iterate through mapping rules (sorted by source start):
   - **Identify gaps:** Parts of the interval not covered by any mapping rule keep their identity (no transformation)
   - **Handle overlaps:** For parts of the interval that overlap with a mapping rule:
     - Calculate the overlap region: `[max(interval.start, rule.source), min(interval.end, rule.source + rule.range))`
     - Translate to destination: `mapped_value = rule.destination + (source_value - rule.source)`
     - Preserve interval length as the mapping is linear
   - Use a cursor to track position within the interval and avoid duplicate processing

4. **Pipeline Propagation:**
   - Start with normalized seed intervals
   - Apply each of the 7 mappings in sequence: seed-to-soil, soil-to-fertilizer, ..., humidity-to-location
   - Normalize intervals after each stage to prevent exponential growth

5. **Extract Result:**
   - After all transformations, examine the resulting location intervals
   - Return the minimum start value among all intervals

**Example Walkthrough:**
~~~
Input seed range [79, 80):
  Step 1: seed-to-soil rule (98→50, range 2) and (50→52, range 48)
          → 79 falls in range [50, 98), maps to 52 + (79-50) = 81 → [81, 82)
  Step 2: soil-to-fertilizer rule (15→0, range 37) and (52→37, range 2) and (0→39, range 15)
          → 81 falls in range [52, 54), maps to 37 + (81-52) = 66 → [66, 67)
  ... (continue through remaining stages)
  Final: location interval [82, 83), minimum = 82
~~~

**Complexity Analysis:**
- **Time:** O(k × (n log n + n × m)) where k is number of mapping stages (7), n is number of intervals, m is number of mapping rules per stage
  - Interval count grows as O(n × m) per stage in worst case
  - Normalization (sorting + merging) keeps it practical
- **Space:** O(n × m × k) for storing intervals at each pipeline stage
- **Practical Performance:** On the full puzzle input, this runs in milliseconds vs. hours for brute force (which would iterate ~600 billion seeds)

**Why This Works:**
1. **Linear Nature of Mappings:** Each rule applies a constant offset: if value A maps to B, then A+1 maps to B+1
2. **Interval Preservation:** The set of intervals that have "never encountered a mapping rule" can be treated as a single interval through multiple stages
3. **Normalization:** Merging reduces interval proliferation, keeping the representation compact

This algorithm demonstrates how understanding the mathematical structure of a problem (linear piecewise functions) can lead to dramatic performance improvements through a change in representation.
