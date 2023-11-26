package agt

import (
	"math"
	"math/rand"
)

type Bernoulli struct {
	P   float64
	Src rand.Source
}

func (b Bernoulli) Rand() float64 {
	var rnd float64
	if b.Src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(b.Src).Float64()
	}
	if rnd < b.P {
		return 1
	}
	return 0
}

type Normal struct {
	Mu    float64 // Mean of the normal distribution
	Sigma float64 // Standard deviation of the normal distribution
	Src   rand.Source
}

func (n Normal) CDF(x float64) float64 {
	return 0.5 * math.Erfc(-(x-n.Mu)/(n.Sigma*math.Sqrt2))
}
