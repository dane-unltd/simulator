package phy

import (
	"math"
	"math/rand"
)

const MaxClusters = 5

type Station struct {
	speed float64
}

type LinkParams struct {
	DelaySpread      float64
	AngularSpreadDep float64
	AngularSpreadArr float64
	ShadowFading     float64
	RicianK          float64
}

type Link struct {
	params LinkParams
	ms, bs *Station
	ch     *Channel
}

type Channel struct {
	RicianK1, RicianK2 float64
	delays             []float64
}

func (link *Link) genChannel() {
	link.ch = &Channel{}
	link.ch.setRicianK(link.params.RicianK)
	link.ch.genDelays(link.params)
}

func (ch *Channel) setRicianK(rk float64) {
	ch.RicianK1 = math.Sqrt(1 / (rk + 1))
	ch.RicianK2 = math.Sqrt(rk / (rk + 1))
}

func (ch *Channel) genDelays(params LinkParams) {
	//TODO: How to get params
	nClusters := 5
	tau := 0.3

	ch.delays = make([]float64, nClusters)

	for i := 0; len(ch.delays); i++ {
		ch.delays[i] = -tau * params.DelaySpread * rand.Float64()
	}
}
