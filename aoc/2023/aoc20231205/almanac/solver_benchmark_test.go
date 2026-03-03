package almanac

import (
	"io"
	"log"
	"sync"
	"testing"
)

type benchmarkDataset struct {
	seedRanges             []SeedRange
	associatedRangesByName map[string][]AlmanacMappingRule
}

var benchmarkDataOnce sync.Once
var benchmarkDataByFile map[string]benchmarkDataset

func loadBenchmarkData() {
	previousWriter := log.Writer()
	log.SetOutput(io.Discard)
	defer log.SetOutput(previousWriter)

	benchmarkDataByFile = make(map[string]benchmarkDataset)

	for _, filename := range []string{"../data/input_test.txt", "../data/input.txt"} {
		seedData, associatedRangesByName := extractData(filename)
		benchmarkDataByFile[filename] = benchmarkDataset{
			seedRanges:             seedRangesFromSeedData(seedData),
			associatedRangesByName: associatedRangesByName,
		}
	}
}

func seedRangesToMinLocationBruteForce(seedRanges []SeedRange, associatedRangesByName map[string][]AlmanacMappingRule) uint64 {
	var minLocation uint64 = ^uint64(0)
	for _, seedRange := range seedRanges {
		for seed := range seedRange.Seeds() {
			soil := seedToSoil(seed, associatedRangesByName)
			fertilizer := soilToFertilizer(soil, associatedRangesByName)
			water := fertilizerToWater(fertilizer, associatedRangesByName)
			light := waterToLight(water, associatedRangesByName)
			temperature := lightToTemperature(light, associatedRangesByName)
			humidity := temperatureToHumidity(temperature, associatedRangesByName)
			location := humidityToLocation(humidity, associatedRangesByName)
			if location < minLocation {
				minLocation = location
			}
		}
	}
	return minLocation
}

func benchmarkPart2Solver(b *testing.B, filename string, solver func([]SeedRange, map[string][]AlmanacMappingRule) uint64) {
	benchmarkDataOnce.Do(loadBenchmarkData)
	dataset := benchmarkDataByFile[filename]

	b.ResetTimer()
	for range b.N {
		_ = solver(dataset.seedRanges, dataset.associatedRangesByName)
	}
}

func BenchmarkPart2Optimized_InputTest(b *testing.B) {
	benchmarkPart2Solver(b, "../data/input_test.txt", seedRangesToMinLocation)
}

func BenchmarkPart2BruteForce_InputTest(b *testing.B) {
	benchmarkPart2Solver(b, "../data/input_test.txt", seedRangesToMinLocationBruteForce)
}

func BenchmarkPart2Optimized_Input(b *testing.B) {
	benchmarkPart2Solver(b, "../data/input.txt", seedRangesToMinLocation)
}

func BenchmarkPart2BruteForce_Input(b *testing.B) {
	benchmarkPart2Solver(b, "../data/input.txt", seedRangesToMinLocationBruteForce)
}
