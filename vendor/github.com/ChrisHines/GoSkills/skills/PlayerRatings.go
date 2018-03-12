package skills

type PlayerRatings map[interface{}]Rating

type RatingAccumulator func(r Rating, a float64) float64

func (pr PlayerRatings) Accum(f RatingAccumulator) (a float64) {
	for _, r := range pr {
		a = f(r, a)
	}
	return
}
