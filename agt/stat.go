package agt

import (
	"math"
	"math/rand"
)

type Normal struct {
	Mu    float64 // Mean of the normal distribution
	Sigma float64 // Standard deviation of the normal distribution
	Src   rand.Source
}

func (n Normal) CDF(x float64) float64 {
	return 0.5 * math.Erfc(-(x-n.Mu)/(n.Sigma*math.Sqrt2))
}
