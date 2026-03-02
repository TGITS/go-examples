package almanac

// AlmanacMappingRule describes one mapping rule in the Almanac.
//
// A value in [Source, Source+Range) maps to the destination interval
// [Destination, Destination+Range) while preserving its offset.
type AlmanacMappingRule struct {
	Source      uint64
	Destination uint64
	Range       uint64
}

// In reports whether value belongs to the source interval of the range.
func (ar *AlmanacMappingRule) In(value uint64) bool {
	return value >= ar.Source && value < ar.Source+ar.Range
}

// GetAssociatedValue converts value to its mapped destination value.
//
// The second return value is true when value is inside the source interval,
// otherwise it returns (0, false).
func (ar *AlmanacMappingRule) GetAssociatedValue(value uint64) (uint64, bool) {
	if ar.In(value) {
		return ar.Destination + (value - ar.Source), true
	}
	return 0, false
}
