package numerics

import (
	"fmt"
	"math"
)

// A range (closed interval) of integers
type Range struct {
	min int
	max int
}

// Construct a range (closed interval).
func NewRange(min, max int) Range {
	if min > max {
		panic(fmt.Errorf("min %v > max %v", min, max))
	}
	return Range{min, max}
}

// Construct a range with a minimum value
func AtLeast(x int) Range {
	return NewRange(x, math.MaxInt32)
}

// Construct a range covering a single integer
func Exactly(x int) Range {
	return NewRange(x, x)
}

func (r *Range) In(x int) bool {
	return r.min <= x && x <= r.max
}
