package almanac

import (
	"reflect"
	"testing"
)

func TestSeeds(t *testing.T) {
	sr := SeedRange{Source: 100, Range: 5}
	expectedSeeds := []uint64{100, 101, 102, 103, 104}

	var actualSeeds []uint64
	for seed := range sr.Seeds() {
		actualSeeds = append(actualSeeds, seed)
	}

	if !reflect.DeepEqual(actualSeeds, expectedSeeds) {
		t.Errorf("Expected seeds %v, but got %v", expectedSeeds, actualSeeds)
	}
}
