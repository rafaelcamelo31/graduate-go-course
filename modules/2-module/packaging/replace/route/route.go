package route

// Replaced by local math package
import "github.com/rafaelcamelo31/graduate-go-course/2-module/packaging/math"

type route struct {
	pointA int
	pointB int
}

func NewRoute(pointA int, pointB int) route {
	return route{pointA: pointA, pointB: pointB}
}

func (r *route) GetRoute() int {
	m := math.NewMath(r.pointA, r.pointB)
	return m.Add()
}
