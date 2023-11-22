package agt

import "math/rand"

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
