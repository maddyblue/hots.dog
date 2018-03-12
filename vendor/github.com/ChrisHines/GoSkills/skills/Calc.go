package skills

const (
	None          = 0x00
	PartialPlay   = 0x01
	PartialUpdate = 0x02
)

const (
	Win  = 1
	Draw = 0
	Lose = -1
)

// Methods required to calculate skills.
type Calc interface {
	// Calculates new ratings based on the prior ratings and team ranks;
	// use 1 for first place, repeat the number for a tie (e.g. 1, 2, 2).
	CalcNewRatings(gi *GameInfo, priors []Team, teamRanks ...int) PlayerRatings

	// Calculates the match quality as the likelihood of all teams
	// drawing (0% = bad, 100% = well matched).
	CalcMatchQual(gi *GameInfo, teams []Team) float64
}
