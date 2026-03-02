package boatrace

import "testing"

func TestNumberOfWaysToWin_1(t *testing.T) {
	var duration uint64 = 7
	var bestDistance uint64 = 9
	verifyNumberOfWaysToWin(t, duration, bestDistance, 4)
}

func TestNumberOfWaysToWin_2(t *testing.T) {
	var duration uint64 = 15
	var bestDistance uint64 = 40
	verifyNumberOfWaysToWin(t, duration, bestDistance, 8)
}

func TestNumberOfWaysToWin_3(t *testing.T) {
	var duration uint64 = 30
	var bestDistance uint64 = 200
	verifyNumberOfWaysToWin(t, duration, bestDistance, 9)
}

func verifyNumberOfWaysToWin(t *testing.T, duration uint64, bestDistance uint64, expectedWays uint64) {
	numWays := NumberOfWaysToWin(duration, bestDistance)

	if numWays != expectedWays {
		t.Errorf("Expected %d ways to win, but got %d", expectedWays, numWays)
	}
}

func TestRecordBreakingsProducts(t *testing.T) {
	expectedProduct := uint64(288)
	actualProduct := RecordBreakingsProducts("../data/input_test.txt")
	if actualProduct != expectedProduct {
		t.Errorf("Expected product %d, but got %d", expectedProduct, actualProduct)
	}
}

func TestRecordBreaking(t *testing.T) {
	expectedWays := uint64(71503)
	actualWays := RecordBreaking("../data/input_test.txt")
	if actualWays != expectedWays {
		t.Errorf("Expected ways %d, but got %d", expectedWays, actualWays)
	}
}
