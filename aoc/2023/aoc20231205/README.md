# AoC 2023 Day 5

## Overview

Go solution for Advent of Code 2023 - Day 5.

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
