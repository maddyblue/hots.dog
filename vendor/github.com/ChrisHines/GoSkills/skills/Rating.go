package skills

import (
	"fmt"
	"github.com/ChrisHines/GoSkills/skills/numerics"
)

type Rating struct {
	mean   float64
	stddev float64
}

func NewRating(mean, stddev float64) Rating {
	return Rating{mean, stddev}
}

func (r Rating) Mean() float64 {
	return r.mean
}

func (r Rating) Stddev() float64 {
	return r.stddev
}

func (r Rating) Variance() float64 {
	return numerics.Sqr(r.stddev)
}

func (r Rating) String() string {
	return fmt.Sprintf("{μ:%.6g σ:%.6g}", r.mean, r.stddev)
}

func MeanSum(r Rating, a float64) float64 {
	return a + r.mean
}

func VarianceSum(r Rating, a float64) float64 {
	return a + r.Variance()
}
