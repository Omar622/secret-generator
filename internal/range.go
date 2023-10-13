package internal

import (
	"fmt"
	"math/rand"
)

// Range is Range struct
type Range struct {
	Start uint
	End   uint
}

// Compare compare between two Ranges
func (r *Range) Compare(other Range) int {
	if r.Start < other.Start {
		return -1
	} else if r.Start == other.Start {
		return 0
	}
	return 2
}

// Count get the number in the range
func (r *Range) Count() uint {
	return r.End - r.Start + 1
}

// String convert Range to string
func (r *Range) String() string {
	return fmt.Sprintf("(%v : %v)", r.Start, r.End)
}

// Shuffle shuffle slice of type Range
func Shuffle(ranges []Range, rnd *rand.Rand) []Range {
	length := int32(len(ranges))
	result := []Range{}

	for possible := length; possible > 0; possible-- {
		i := rnd.Int31n(possible)
		result = append(result, ranges[i])
		ranges[i], ranges[possible-1] = ranges[possible-1], ranges[i]
	}

	return result
}
