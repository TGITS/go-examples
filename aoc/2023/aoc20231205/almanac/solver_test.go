package almanac

import "testing"

func TestGetMinimumLocationForPart1(t *testing.T) {
	expectedMinLocation := uint64(35)
	minLocation := GetMinimumLocationForPart1("../data/input_test.txt")
	if minLocation != expectedMinLocation {
		t.Errorf("Expected minimum location %d, got %d", expectedMinLocation, minLocation)
	}
}

func TestGetMinimumLocationForPart2(t *testing.T) {
	expectedMinLocation := uint64(46)
	minLocation := GetMinimumLocationForPart2("../data/input_test.txt")
	if minLocation != expectedMinLocation {
		t.Errorf("Expected minimum location %d, got %d", expectedMinLocation, minLocation)
	}
}
