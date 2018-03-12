package skills

const (
	defaultInitialMean     = 25.0
	defaultDrawProbability = 0.10
	defaultInitialStddev   = defaultInitialMean / 3.0
	defaultBeta            = defaultInitialMean / 6.0
	defaultDynamicsFactor  = defaultInitialMean / 300.0
)

type GameInfo struct {
	InitialMean     float64
	DrawProbability float64
	InitialStddev   float64
	Beta            float64
	DynamicsFactor  float64
}

func (this *GameInfo) DefaultRating() Rating {
	return NewRating(this.InitialMean, this.InitialStddev)
}

var DefaultGameInfo = &GameInfo{
	InitialMean:     defaultInitialMean,
	DrawProbability: defaultDrawProbability,
	InitialStddev:   defaultInitialStddev,
	Beta:            defaultBeta,
	DynamicsFactor:  defaultDynamicsFactor,
}
