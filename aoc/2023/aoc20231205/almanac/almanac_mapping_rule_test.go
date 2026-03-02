package almanac

import "testing"

func TestIn(t *testing.T) {
	ar := AlmanacMappingRule{Source: 10, Destination: 20, Range: 5}
	if !ar.In(12) {
		t.Errorf("Expected 12 to be in range %v", ar)
	}
	if ar.In(9) {
		t.Errorf("Expected 9 to not be in range %v", ar)
	}
}

func TestGetAssociatedValue(t *testing.T) {
	ar := AlmanacMappingRule{Source: 10, Destination: 20, Range: 5}
	value, ok := ar.GetAssociatedValue(12)
	if !ok {
		t.Errorf("Expected 12 to be in range %v", ar)
	} else if value != 22 {
		t.Errorf("Expected associated value for 12 to be 22, got %d", value)
	}
}
