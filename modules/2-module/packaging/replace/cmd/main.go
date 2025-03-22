package main

import (
	"log"

	"github.com/rafaelcamelo31/graduate-go-course/2-module/packaging/math"
	"github.com/rafaelcamelo31/graduate-go-course/2-module/packaging/route"
)

/*
   go.mod => is for managing dependencies
   go.sum => is for tracking the exact version of dependencies
   Indirect dependencies are dependencies of dependencies and have // indirect at the end of the line

   Using replace has its own limitations:
   - It only works for local packages
   - It is not a good practice to use replace in a production environment
   - It is not a good practice to use replace in a shared repository
*/

func main() {
	m := math.NewMath(1, 2)
	log.Println("Math:", m.Add())

	r := route.NewRoute(1, 2)
	log.Println("Route:", r.GetRoute())
}
