package gl

import (
	"guillermoSb/glRayTracing/numg"
	"math"
)



type sphere struct {
	center numg.V3
	radius float64
	material material
}

func NewSphere(center numg.V3, radius float64, material material) *sphere {
	s := sphere{center: center, radius: radius, material: material}
	return &s
}


func (sphere *sphere) rayIntersect(origin,dir numg.V3) *intersect{
	l := numg.Subtract(sphere.center, origin)
	tca := numg.V3DotProduct(l, dir)
	d := math.Sqrt(math.Pow(l.Magnitude(),2) - math.Pow(tca,2))
	if d > sphere.radius {
		return nil
	}
	thc := math.Sqrt(math.Pow(sphere.radius,2) - math.Pow(d,2))
	t0 := tca - thc
	t1 := tca + thc

	// ? Por que 
	if t0 < 0 {
		t0 = t1
	}
	if t0 < 0 {
		return nil
	}
	p := numg.Add(origin, numg.MultiplyVectorWithConstant(numg.NormalizeV3(dir), t0))
	return NewIntersect(t0, p, numg.NormalizeV3(numg.Subtract(p, sphere.center)), struct{material}{
		material: sphere.material,
	})
}