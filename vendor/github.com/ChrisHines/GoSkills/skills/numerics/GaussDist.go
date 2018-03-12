package numerics

import (
	"fmt"
	"math"
)

const logSqrt2Pi = 0.91893853320467274178032973640562

func GaussAt(x float64) float64 {
	return math.Exp(-x*x/2) / (math.Sqrt2 * math.SqrtPi)
}

func GaussCumulativeTo(x float64) float64 {
	return math.Erf(x/math.Sqrt2)/2 + 0.5
}

func GaussInvCumulativeTo(x, mean, stddev float64) float64 {
	// From numerical recipes, page 320
	return mean - math.Sqrt(2)*stddev*InvErfc(2*x)
}

// Inverse of complementary error function. Returns x such that erfc(x) = p for argument p.
func InvErfc(p float64) float64 {
	// From page 265 of numerical recipes
	if p >= 2.0 {
		return -100
	}
	if p <= 0.0 {
		return 100
	}

	var pp float64
	if p < 1.0 {
		pp = p
	} else {
		pp = 2 - p
	}

	t := math.Sqrt(-2 * math.Log(pp/2.0)) // Initial guess
	x := -0.70711 * ((2.30753+t*0.27061)/(1.0+t*(0.99229+t*0.04481)) - t)

	for j := 0; j < 2; j++ {
		err := math.Erfc(x) - pp
		x += err / (1.12837916709551257*math.Exp(-(x*x)) - x*err) // Halley
	}

	if p < 1.0 {
		return x
	}
	return -x
}

type GaussDist struct {
	Mean          float64
	Stddev        float64
	Precision     float64
	PrecisionMean float64
	Variance      float64
}

func NewGaussDist(mean, stddev float64) *GaussDist {
	variance := stddev * stddev
	precision := 1 / variance
	return &GaussDist{
		Mean:          mean,
		Stddev:        stddev,
		Variance:      variance,
		Precision:     precision,
		PrecisionMean: precision * mean,
	}
}

func (z *GaussDist) String() string {
	return fmt.Sprintf("{μ:%.6g σ:%.6g}", z.Mean, z.Stddev)
}

// Sub sets z to the difference x-y and returns z.
func (z *GaussDist) Sub(x, y *GaussDist) *GaussDist {
	z.Mean = x.Mean - y.Mean
	z.Variance = x.Variance + y.Variance
	z.Stddev = math.Sqrt(z.Variance)
	z.Precision = 1 / z.Variance
	z.PrecisionMean = z.Mean * z.Precision
	return z
}

// Mul sets z to the product x*y and returns z.
func (z *GaussDist) Mul(x, y *GaussDist) *GaussDist {
	z.Precision = x.Precision + y.Precision
	z.PrecisionMean = x.PrecisionMean + y.PrecisionMean
	z.fromPrecisionMean()
	return z
}

// Div sets z to the product x/y and returns z.
func (z *GaussDist) Div(x, y *GaussDist) *GaussDist {
	z.Precision = x.Precision - y.Precision
	z.PrecisionMean = x.PrecisionMean - y.PrecisionMean
	z.fromPrecisionMean()
	return z
}

// CumulativeTo returns the cumulative distrubution function evaluated at x.
func (z *GaussDist) CumulativeTo(x float64) float64 {
	return GaussCumulativeTo((x - z.Mean) / z.Stddev)
}

func (z *GaussDist) fromPrecisionMean() {
	z.Variance = 1 / z.Precision
	z.Stddev = math.Sqrt(z.Variance)
	z.Mean = z.PrecisionMean / z.Precision
}

// Returns the log product normalization of x and y.
func LogProdNorm(x, y *GaussDist) float64 {
	if x.Precision == 0 || y.Precision == 0 {
		return 0
	}

	varSum := x.Variance + y.Variance
	meanDiff := x.Mean - y.Mean
	meanDiff2 := meanDiff * meanDiff

	return -logSqrt2Pi - (math.Log(varSum)+meanDiff2/varSum)/2.0
}

// Returns the log ratio normalization of x and y.
func LogRatioNorm(x, y *GaussDist) float64 {
	if x.Precision == 0 || y.Precision == 0 {
		return 0
	}

	varDiff := y.Variance - x.Variance
	meanDiff := x.Mean - y.Mean
	meanDiff2 := meanDiff * meanDiff

	return math.Log(y.Variance) + logSqrt2Pi - math.Log(varDiff)/2 + meanDiff2/(2*varDiff)
}

// Computes the absolute difference between two Gaussians
func AbsDiff(x, y *GaussDist) float64 {
	return math.Max(math.Abs(x.PrecisionMean-y.PrecisionMean), math.Sqrt(math.Abs(x.Precision-y.Precision)))
}
