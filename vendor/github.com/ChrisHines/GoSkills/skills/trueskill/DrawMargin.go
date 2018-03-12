package trueskill

import (
	"github.com/ChrisHines/GoSkills/skills/numerics"
	"math"
)

func drawMarginFromDrawProbability(drawProbability, beta float64) float64 {
	// Derived from TrueSkill technical report (MSR-TR-2006-80), page 6
	//
	// draw probability = 2 * CDF(margin/(sqrt(n1+n2)*beta)) -1
	//
	// implies
	//
	// margin = inversecdf((draw probability + 1)/2) * sqrt(n1+n2) * beta
	// n1 and n2 are the number of players on each team
	return numerics.GaussInvCumulativeTo((drawProbability+1)/2, 0, 1) * math.Sqrt(1+1) * beta
}
