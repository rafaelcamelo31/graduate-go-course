package main

import (
	"log"

	"github.com/rafaelcamelo31/math"
	"github.com/rafaelcamelo31/route"
)

/*
   Using workspace centralizes the dependencies in a single place
   and makes it easier to manage them.
   Use go mod tidy -e to update depencies, ignoring not found ones
*/

func main() {
	m := math.NewMath(1, 2)
	log.Println("Math:", m.Add())

	r := route.NewRoute(1, 2)
	log.Println("Route:", r.GetRoute())
}
