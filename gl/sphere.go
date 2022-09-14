package gl

import (
	"guillermoSb/glRayTracing/numg"
	"math"
)



type sphere struct {
	center numg.V3
	radius float64
}

func NewSphere(center numg.V3, radius float64) *sphere {
	s := sphere{center: center, radius: radius}
	return &s
}


func (sphere *sphere) rayIntersect(origin,dir numg.V3) bool{
	l := numg.Subtract(sphere.center, origin)
	tca := numg.V3DotProduct(l, dir)
	d := math.Sqrt(math.Pow(l.Magnitude(),2) - math.Pow(tca,2))
	if d > sphere.radius {
		return false
	}
	return true
}