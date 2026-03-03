package boatrace

import (
	"io"
	"log"
	"sync"
	"testing"
)

var benchmarkDataOnce sync.Once
var benchmarkPart1TestInput uint64
var benchmarkPart2TestInput uint64

func loadBenchmarkData() {
	previousWriter := log.Writer()
	log.SetOutput(io.Discard)
	defer log.SetOutput(previousWriter)

	benchmarkPart1TestInput = RecordBreakingsProducts("../data/input_test.txt")
	benchmarkPart2TestInput = RecordBreaking("../data/input_test.txt")
}

func BenchmarkNumberOfWaysToWin_SmallRace(b *testing.B) {
	var duration uint64 = 30
	var bestDistance uint64 = 200

	b.ResetTimer()
	for range b.N {
		_ = NumberOfWaysToWin(duration, bestDistance)
	}
}

func BenchmarkRecordBreakingsProducts_InputTest(b *testing.B) {
	benchmarkDataOnce.Do(loadBenchmarkData)
	_ = benchmarkPart1TestInput

	previousWriter := log.Writer()
	log.SetOutput(io.Discard)
	defer log.SetOutput(previousWriter)

	b.ResetTimer()
	for range b.N {
		_ = RecordBreakingsProducts("../data/input_test.txt")
	}
}

func BenchmarkRecordBreaking_InputTest(b *testing.B) {
	benchmarkDataOnce.Do(loadBenchmarkData)
	_ = benchmarkPart2TestInput

	previousWriter := log.Writer()
	log.SetOutput(io.Discard)
	defer log.SetOutput(previousWriter)

	b.ResetTimer()
	for range b.N {
		_ = RecordBreaking("../data/input_test.txt")
	}
}
