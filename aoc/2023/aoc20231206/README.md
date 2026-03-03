# AoC 2023 Day 6

**French version available:** [README_FR.md](README_FR.md)

## Overview

Go solution for [Advent of Code 2023 - Day 6](https://adventofcode.com/2023/day/6).

## Run

From this folder (`aoc/2023/aoc20231206`):

```powershell
go run .
```

## Test

From this folder (`aoc/2023/aoc20231206`):

```powershell
go test ./...
```

## Benchmarks

From this folder (`aoc/2023/aoc20231206`):

```powershell
# Run all boatrace benchmarks
go test ./boatrace -bench . -benchmem -run "^$"

# One-shot examples
go test ./boatrace -bench "BenchmarkNumberOfWaysToWin_SmallRace$" -benchmem -benchtime=1x -run "^$"
go test ./boatrace -bench "BenchmarkRecord(Breaking|BreakingsProducts)_InputTest$" -benchmem -benchtime=1x -run "^$"
```

Or run the helper scripts:

```powershell
powershell -ExecutionPolicy Bypass -File ".\run_benchmarks.ps1"
```

```bash
bash ./run_benchmarks.sh
```

## Input files

- `data/input_test.txt`
- `data/input.txt`
  - This file is specific to each AoC user and is not committed to the repository: you must provide your own input file.
