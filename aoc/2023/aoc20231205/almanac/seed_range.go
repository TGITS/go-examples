package almanac

import "iter"

// SeedRange represents a contiguous range of seed IDs.
type SeedRange struct {
	Source uint64
	Range  uint64
}

// Seeds returns an iterator over all seed IDs contained in the range.
//
// The generated sequence is [Source, Source+Range).
func (sr *SeedRange) Seeds() iter.Seq[uint64] {
	return func(yield func(uint64) bool) {
		for i := uint64(0); i < sr.Range; i++ {
			if !yield(sr.Source + i) {
				return
			}
		}
	}
}
