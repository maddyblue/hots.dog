package trueskill

import (
	"github.com/ChrisHines/GoSkills/skills/numerics"
	"math"
)

// These functions from the bottom of page 4 of the TrueSkill paper.

// The "V" function where the team performance difference is greater than the draw margin.
// In the reference F# implementation, this is referred to as "the additive
// correction of a single-sided truncated Gaussian with unit variance."
// In the paper drawMargin is referred to as just "ε".
func vExceedsMarginC(perfDiff, drawMargin, c float64) float64 {
	return vExceedsMargin(perfDiff/c, drawMargin/c)
}

func vExceedsMargin(perfDiff, drawMargin float64) float64 {
	denom := numerics.GaussCumulativeTo(perfDiff - drawMargin)
	if denom < 2.222758749e-162 {
		return -perfDiff + drawMargin
	}
	return numerics.GaussAt(perfDiff-drawMargin) / denom
}

// The "W" function where the team performance difference is greater than the draw margin.
// In the reference F# implementation, this is referred to as "the multiplicative
// correction of a single-sided truncated Gaussian with unit variance."
func wExceedsMarginC(perfDiff, drawMargin, c float64) float64 {
	return wExceedsMargin(perfDiff/c, drawMargin/c)
}

func wExceedsMargin(perfDiff, drawMargin float64) float64 {
	denom := numerics.GaussCumulativeTo(perfDiff - drawMargin)
	if denom < 2.222758749e-162 {
		if perfDiff < 0.0 {
			return 1.0
		}
		return 0.0
	}

	vWin := vExceedsMargin(perfDiff, drawMargin)
	return vWin * (vWin + perfDiff - drawMargin)
}

// the additive correction of a double-sided truncated Gaussian with unit variance
func vWithinMarginC(perfDiff, drawMargin, c float64) float64 {
	return vWithinMargin(perfDiff/c, drawMargin/c)
}

// from F#:
func vWithinMargin(perfDiff, drawMargin float64) float64 {
	perfDiffAbs := math.Abs(perfDiff)
	denom := numerics.GaussCumulativeTo(drawMargin-perfDiffAbs) - numerics.GaussCumulativeTo(-drawMargin-perfDiffAbs)
	if denom < 2.222758749e-162 {
		if perfDiff < 0.0 {
			return -perfDiff - drawMargin
		}
		return -perfDiff + drawMargin
	}

	numerator := numerics.GaussAt(-drawMargin-perfDiffAbs) - numerics.GaussAt(drawMargin-perfDiffAbs)
	if perfDiff < 0.0 {
		return -numerator / denom
	}
	return numerator / denom
}

// the multiplicative correction of a double-sided truncated Gaussian with unit variance
func wWithinMarginC(perfDiff, drawMargin, c float64) float64 {
	return wWithinMargin(perfDiff/c, drawMargin/c)
}

// From F#:
func wWithinMargin(perfDiff, drawMargin float64) float64 {
	perfDiffAbs := math.Abs(perfDiff)
	denom := numerics.GaussCumulativeTo(drawMargin-perfDiffAbs) - numerics.GaussCumulativeTo(-drawMargin-perfDiffAbs)

	if denom < 2.222758749e-162 {
		return 1.0
	}
	vt := vWithinMargin(perfDiffAbs, drawMargin)

	return vt*vt + ((drawMargin-perfDiffAbs)*numerics.GaussAt(drawMargin-perfDiffAbs)-(-drawMargin-perfDiffAbs)*numerics.GaussAt(-drawMargin-perfDiffAbs))/denom
}
