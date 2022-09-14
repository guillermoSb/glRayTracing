package gl

import "guillermoSb/glRayTracing/numg"



type sphere struct {
	center numg.V3
	radius float64
}

func NewSphere(center numg.V3, radius float64) *sphere {
	s := sphere{center: center, radius: radius}
	return &s
}


func (sphere *sphere) rayIntersect(origin,dir numg.V3) bool{
	return true
}