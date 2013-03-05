package main

import (
	"fmt"
	"github.com/skelterjohn/go.matrix"
	"simulation/channels"
)

func main() {
	nUsers := 6
	simulationTime := 400
	nTx := 4
	timeScale := 50
	speed := 40

	a := matrix.Eye(5)

	fmt.Println(a)

	simTime := 100

	chans := channels.flatFadingMIMO(nTx, 1, nUsers, simTime)

}
